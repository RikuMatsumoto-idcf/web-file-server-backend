# web-file-server-backend
goの勉強用のサンプルウェブアプリ（Echo + Clean Architecture版）

## ⚠️ 重要：学習用スケルトンコード

このプロジェクトは **学習用の骨組み** です。ほとんどのメソッドは `panic("implement me")` となっており、実装はあなた自身が行う必要があります。

実装ガイドは [docs/implementation_guide.md](docs/implementation_guide.md) を参照してください。

## プロジェクト構成（Clean Architecture）

```
.
├── cmd/
│   └── server/          # アプリケーションエントリーポイント
│       └── main.go      # Echoサーバーの起動とDI
├── internal/
│   ├── domain/          # ドメイン層（ビジネスロジックの中心）
│   │   ├── file_model.go    # Fileエンティティ、値オブジェクト
│   │   └── errors.go        # ドメインエラー定義
│   ├── usecase/         # ユースケース層（アプリケーションのビジネスルール）
│   │   └── file_usecase.go  # ファイル操作のビジネスロジック
│   ├── infrastructure/  # インフラ層（外部システムとのやり取り）
│   │   └── file_repository.go  # ファイルストレージの実装
│   └── handler/         # ハンドラ層（HTTPリクエスト/レスポンス）
│       └── file_handler.go     # Echoハンドラー
├── docs/
│   ├── implementation_guide.md  # 📖 実装ガイド（必読）
│   └── api/
│       └── openapi.yaml         # API仕様書
├── go.mod
└── README.md
```

### レイヤー間の依存関係

```
Handler → Usecase → Domain
Infrastructure → Domain, Usecase
```

内側のレイヤー（Domain）は外側のレイヤー（Handler, Infrastructure）に依存しません。

## 使用技術

- **Go 1.24+**
- **Echo v4**: 高性能なWebフレームワーク
- **Clean Architecture**: レイヤー分離による保守性の高い設計

## 開発の始め方

### 1. 実装ガイドを読む

まず [docs/implementation_guide.md](docs/implementation_guide.md) を読んでください。以下の内容が含まれています：

- Clean Architectureの概要と依存関係の図
- 各レイヤーの役割と責任
- 実装手順（Step by Step）
- Echoフレームワークの使い方
- 依存性注入(DI)の説明

### 2. 依存関係のインストール

```bash
go mod download
```

### 3. 実装を開始

推奨する実装順序：

1. **Domain層**: `internal/domain/file_model.go` の `NewFileName` を実装
2. **Infrastructure層**: `internal/infrastructure/file_repository.go` の各メソッドを実装
3. **Usecase層**: `internal/usecase/file_usecase.go` の各メソッドを実装
4. **Handler層**: `internal/handler/file_handler.go` の各メソッドを実装

### 4. サーバーの起動

```bash
go run cmd/server/main.go
```

サーバーは `http://localhost:8080` で起動します。


## APIエンドポイント（実装後に使用可能）

実装が完了すると、以下のエンドポイントが使用できるようになります。

### ファイル一覧
- **URL**: `/api/files`
- **Method**: GET
- **Success**: `200 OK` + JSON array

例:
```bash
curl -i http://localhost:8080/api/files
```

### ファイルアップロード
- **URL**: `/api/files/:name`
- **Method**: PUT
- **Body**: raw bytes（multipartではない）
- **Success**: `204 No Content`

例:
```bash
printf 'hello' | curl -i -X PUT --data-binary @- http://localhost:8080/api/files/hello.txt
```

### ファイルダウンロード
- **URL**: `/api/files/:name`
- **Method**: GET
- **Success**: `200 OK` + `Content-Type: application/octet-stream` + raw bytes
- **Not Found**: `404`

例:
```bash
curl -i http://localhost:8080/api/files/hello.txt
```

### ファイル削除
- **URL**: `/api/files/:name`
- **Method**: DELETE
- **Success**: `204 No Content`
- **Not Found**: `404`

例:
```bash
curl -i -X DELETE http://localhost:8080/api/files/hello.txt
```

## テスト

```bash
# 全てのテストを実行
go test ./...

# 特定のパッケージをテスト
go test ./internal/domain

# カバレッジ付きで実行
go test -cover ./...
```

## 学習のヒント

- 各ファイルには詳細なTODOコメントがあります
- `panic("implement me")` を1つずつ実装していってください
- テストを書きながら実装すると理解が深まります
- 困ったときは [docs/implementation_guide.md](docs/implementation_guide.md) を参照してください

## Dev Container（オプション）

このプロジェクトはDev Container環境でも動作します。PostgreSQLデータベースを使用したい場合は、Dev Containerを利用してください。

詳細は `.devcontainer/` ディレクトリを参照してください。
