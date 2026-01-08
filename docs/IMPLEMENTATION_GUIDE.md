# 実装順序ガイド（Go初学者向け / OpenAPI準拠 / Clean Architecture）

対象API仕様: `docs/api/openapi.yaml`

このガイドのゴール:
- `GET /api/health` が JSON を返す
- `PUT /api/files/{name}` が raw bytes を保存して `204` を返す
- `GET /api/files/{name}` が raw bytes を返す（無ければ `404`）
- 依存方向を「外側(HTTP/FS) → 内側(usecase/domain)」にする

前提:
- 既存コード（`internal/api`）は、動作確認できるまで残してOK
- ルーティングは Go 1.22 の `http.ServeMux` を使う想定（`/api/files/{name}` パターン）

---

## 進め方（大原則）

- 1ステップごとに `go test ./...` を通す
- 先に内側（domain → port → usecase）を固め、最後に外側（http → cmd/server）を繋ぐ
- エラーは「domain/usecaseでは種類だけ決める」「HTTPでステータスやJSONに変換する」

---

## Step 0: 現状確認（壊して良い前提でも、今の状態を知る）

やること:
- `go test ./...` が通るか確認
- `cmd/server/main.go` のルートが OpenAPI と違うことを把握（今は `/api/files`）

合格条件:
- テストが現状で通る（通らないなら、このガイドのStep 8以降で更新していく）

---

## Step 1: domain（名前バリデーションとエラー種類を確定）

作るもの（例）:
- `internal/domain/file.go`
  - `type FileName string`
  - `func NewFileName(raw string) (FileName, error)`
  - ルール例（最小でOK）:
    - 空文字NG
    - `/` と `\\` を含むのはNG（ディレクトリを作れないように）
    - `..` を含むのはNG（パストラバーサル対策）
- `internal/domain/errors.go`
  - `var ErrInvalidName = errors.New("invalid name")`
  - `var ErrNotFound = errors.New("not found")`
  - `var ErrAlreadyExists = errors.New("already exists")`
  - （必要なら）`ErrTooLarge`

テスト:
- `internal/domain/file_test.go`
  - OKケース: `"a.txt"`, `"photo.png"`
  - NGケース: `""`, `"../a"`, `"a/b"`, `"a\\b"`

合格条件:
- `go test ./...` が通る

---

## Step 2: port（ストレージ抽象を決める）

作るもの:
- `internal/port/filestore.go`
  - `type FileStore interface {`
    - `Put(ctx context.Context, name domain.FileName, r io.Reader) error`
    - `Get(ctx context.Context, name domain.FileName) (io.ReadCloser, error)`
    - `Exists(ctx context.Context, name domain.FileName) (bool, error)`
  - 注: `domain.FileName` を引数に取ると安全（HTTP層の生文字列が内側に入りにくい）

合格条件:
- `go test ./...` が通る（まだ実装がなくてもOK）

---

## Step 3: infra（ローカルFSの FileStore 実装を作る）

（DB保存に変更するため、Step 3 は「DB実装」前提に読み替え）

作るもの:
- `internal/infra/storage/db/filestore.go`
  - `type Store struct { DB *sql.DB }` など
  - `Put/Get/Exists` を実装
  - テーブル例: `files(name PRIMARY KEY, data BYTEA)`
  - `Put` はトランザクションを使う（必要なら）

実装の注意:
- `Get` で存在しない場合は `domain.ErrNotFound` を返すように寄せる
- `Put` 自体は「既にあるなら上書き」でもよい（上書き禁止の判断はusecaseで行う）
- DBの一意制約違反を `domain.ErrAlreadyExists` 相当に寄せたい場合は、ここで変換するか、usecase側の `Exists` と組み合わせる

テスト:
- `internal/infra/storage/db/filestore_test.go`
  - テストDBを用意（Postgres）
  - `Put` → `Exists` → `Get` で内容一致
  - 存在しない `Get` が `domain.ErrNotFound` になる

合格条件:
- `go test ./...` が通る

---

## Step 4: usecase（Put/Get/Health を作る）

作るもの:
- `internal/usecase/health_check.go`
  - `HealthCheck(ctx)` が `{status:"ok", message:"Server is running"}` を返す
- `internal/usecase/put_file.go`
  - 入力: `rawName string`, `body io.Reader`（or `domain.FileName` を受ける）
  - `domain.NewFileName(rawName)`
  - overwrite禁止なら `store.Exists` を見て `domain.ErrAlreadyExists`
  - `store.Put`
- `internal/usecase/get_file.go`
  - `domain.NewFileName(rawName)`
  - `store.Get`

テスト:
- `internal/usecase/*_test.go`
  - `FileStore` のテスト用実装（in-memory）を用意して分岐を確認
  - overwrite禁止時のAlreadyExists
  - NotFound
  - invalid name

合格条件:
- `go test ./...` が通る

---

## Step 5: adapter/http（OpenAPI準拠のHTTPハンドラを作る）

作るもの:
- `internal/adapter/http/error_response.go`
  - `type ErrorResponse struct { Code, Message string }`
  - `writeError(w, status, code, message)`
- `internal/adapter/http/router.go`
  - `func NewMux(deps ...) *http.ServeMux`
  - `GET /api/health`
  - `PUT /api/files/{name}`
  - `GET /api/files/{name}`
- `internal/adapter/http/health_handler.go`
  - `HealthCheck` usecase を呼び、JSONで返す
- `internal/adapter/http/files_handler.go`
  - PUT:
    - `name := r.PathValue("name")`
    - `r.Body` をサイズ制限付きで usecase に渡す
    - 成功 `204`
  - GET:
    - `name := r.PathValue("name")`
    - usecase から `io.ReadCloser` を受け、`io.Copy(w, rc)`
    - `Content-Type: application/octet-stream`

エラーマッピング（例）:
- `domain.ErrInvalidName` → 400 + `{code:"bad_request"}`
- `domain.ErrAlreadyExists` → 409 + `{code:"conflict"}`（code文字列は好みでOK、ただしOpenAPIは例なので固定必須ではない）
- `domain.ErrNotFound` → 404 + `{code:"not_found"}`
- サイズ超過 → 413 + `{code:"payload_too_large"}`

テスト:
- `internal/adapter/http/*_test.go`
  - `httptest.NewServer(NewMux(...))`
  - PUT 204 / GET 200 / GET 404 / PUT 400 / PUT 413

合格条件:
- `go test ./...` が通る

---

## Step 6: cmd/server（DIして起動）

やること:
- `cmd/server/main.go` を「新しい router.NewMux を起動」へ差し替える
- DB接続（`DB_DSN`）や最大サイズなどの設定を読む

合格条件:
- `go run ./cmd/server` で起動できる

手動確認（例）:
- health:
  - `curl -i http://localhost:8080/api/health`
- put/get:
  - `printf 'hello' | curl -i -X PUT --data-binary @- http://localhost:8080/api/files/hello.txt`
  - `curl -i http://localhost:8080/api/files/hello.txt`

---

## Step 7: 既存テストの移行（integrationの位置づけを整理）

やること:
- 既存の `internal/api/*_test.go` と `internal/api/integration_test.go` は、新しいHTTP層に合わせて更新
- ルートが `/api/files/{name}` 前提になるように変更

合格条件:
- `go test ./...` が通る

---

## Step 8: 旧実装の整理（internal/api を卒業）

やること:
- `internal/api` を参照している箇所をゼロにする
- `internal/api` ディレクトリを削除（or 役割が残るなら adapter/http に統合）

合格条件:
- `go test ./...` が通る
- OpenAPIの2エンドポイントが期待どおりに動く

---

## つまずきやすいポイント（先回り）

- `PUT` のボディは multipart ではなく raw bytes（`--data-binary` を使う）
- 413はHTTP層で発生させるのが簡単（`http.MaxBytesReader`）
- `io.ReadCloser` は必ず `defer Close()`（usecase側 or http側で責務を決めて統一）
- `domain.FileName` のルールは「少し厳しめ」くらいが安全（後から緩められる）
