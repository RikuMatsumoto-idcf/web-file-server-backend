# Clean Architecture 実装TODO（OpenAPI準拠）

対象OpenAPI: `docs/api/openapi.yaml`

> 目的: HTTP(外側) / Usecase(内側) / Domain(内側) / Infra(外側) を分離し、依存方向を「外→内」に固定する。
> ここには「各フォルダで実装が必要な処理内容」だけを書き、設計説明は最小限にする。

---

## cmd/server

- DI（依存の組み立て）を行う
  - `FileStore` / `FileRepository` 実装（DB）を生成
  - `usecase` のサービスを生成（`PutFile`, `GetFile`）
  - HTTPハンドラ/ルータに注入して `http.ListenAndServe` で起動
- 設定値を読み込む（環境変数 or デフォルト）
  - 例: `PORT`, `DB_DSN`, `MAX_UPLOAD_BYTES`, `ALLOW_OVERWRITE`

---

## internal/adapter/http（Controller）

### ルーティング
- `PUT /api/files/{name}` を登録（raw bytes）
- `GET /api/files/{name}` を登録（raw bytes）
- Go 1.22 の `http.ServeMux` を使う場合:
  - パターン `"/api/files/{name}"`
  - `r.PathValue("name")` で取得

### リクエスト処理
- `PUT /api/files/{name}`
  - `name` の取得・空チェック
  - ボディを読み取り（ストリーミング前提）
  - サイズ制限を適用（413）
    - 例: `http.MaxBytesReader(w, r.Body, maxBytes)`
  - 成功時: `204 No Content`（ボディ無し）
- `GET /api/files/{name}`
  - `name` の取得・空チェック
  - 成功時: `200 OK` + `Content-Type: application/octet-stream` + ボディにbytes

### エラー応答（OpenAPI `ErrorResponse`）
- JSON: `{ "code": string, "message": string }`
- ユースケース/ドメインのエラーをHTTPにマッピング
  - 400: invalid request / invalid name
  - 404: not found
  - 409: already exists（上書き禁止の場合）
  - 413: payload too large
- 返却時に `Content-Type: application/json`

### テスト（HTTP層）
- ルーティングが正しいこと（`/api/files` ではなく `/api/files/{name}`）
- PUT成功で204・ボディ無し
- GET成功で200・`application/octet-stream`
- 400/404/409/413 のJSON形がOpenAPIどおり

---

## internal/usecase（Application Service）

### PutFile
- `PutFile(ctx, name, reader)` を実装
  - `domain.FileName` に変換/バリデーション
  - `port.FileStore` に保存を委譲
  - `ALLOW_OVERWRITE=false` のとき
    - 既存チェックして存在すれば 409 相当のエラーを返す
  - エラーをドメイン/ユースケースの型に正規化（HTTPを知らない）

### GetFile
- `GetFile(ctx, name)` を実装
  - `domain.FileName` に変換/バリデーション
  - `port.FileStore` から取得
  - 無ければ NotFound 相当のエラーを返す

### テスト（ユースケース層）
- `FileStore` をモック/インメモリにして分岐を確認
  - overwrite禁止時のAlreadyExists
  - NotFound
  - invalid name

---

## internal/domain（Entity / ValueObject / Domain Errors）

### 値オブジェクト
- `FileName` を実装
  - `minLength: 1` を満たす
  - パストラバーサル防止（例: `..`, `/`, `\\` などを禁止する等）
  - ルールに反する場合は `ErrInvalidName`

### ドメイン/ユースケース共通エラー
- 最低限のエラー種別を定義（HTTPを知らない）
  - `ErrInvalidName`
  - `ErrNotFound`
  - `ErrAlreadyExists`
  - `ErrTooLarge`（必要なら）

### テスト（ドメイン層）
- `FileName` の正常/異常ケース

---

## internal/port（Interface / Port）

### ストレージ抽象
- `FileStore` インタフェースを定義
  - `Put(ctx, name, r io.Reader) error`
  - `Get(ctx, name) (io.ReadCloser, error)`
  - `Exists(ctx, name) (bool, error)`（overwrite禁止時に必要）
- 返すべきエラーの契約を決める（例: NotFound/AlreadyExists の表現）

※ 保存先がDBでも、usecaseから見ると「永続化ポート」である点は同じ。
  DB固有の概念（テーブル、SQL、トランザクション）は `infra` 側に閉じ込める。

---

## internal/infra/storage/db（FileStore実装：DB）

- `FileStore` をDBに保存する実装（Postgres想定）
  - テーブル例: `files(id PRIMARY KEY, name UNIQUE, data BYTEA, created_at, updated_at)`
  - `id` を主キーにし、`name` は一意制約で担保
  - `data` にバイナリ（Postgres: `BYTEA`）を保存
  - Exists / Get / Put を実装（SQL/ORMはここに閉じ込める）
- 返すエラーを `domain/usecase` の想定に合わせて正規化

### DBスキーマ（Dev Containerのinitdb）
- `.devcontainer/initdb/` に `files` テーブル作成SQLを追加
  - 例: `03_files.sql`
  - `CREATE TABLE IF NOT EXISTS files (...);`

### テスト（infra）
- テストDB（Postgres）で Put/Get/Exists の基本動作
  - Dev Containerの `db` サービスを使う
  - DSN例: `postgres://user:password@db:5432/db?sslmode=disable`

---

## internal/infra/config

- 設定読み込み（環境変数 + デフォルト）
  - `PORT`（default: 8080）
  - `DB_DSN`（例: `postgres://user:pass@localhost:5432/dbname?sslmode=disable`）
  - `MAX_UPLOAD_BYTES`（default: 適当な上限）
  - `ALLOW_OVERWRITE`（default: false 推奨）

---

## 既存コードの移行メモ（作業項目）

- `internal/api/*` は段階的に置き換え（最終的に削除/リネーム）
- 既存の `GET /api/files` はOpenAPIに無いので削除（もしくは別APIとして別specに）
- 既存テスト `internal/api/*_test.go` と `internal/api/integration_test.go` は、新構成に合わせて移設・更新
