# web-file-server-backend
goの勉強用のサンプルウェブアプリ

## プロジェクト構成

```
.
├── cmd/
│   └── server/          # アプリケーションエントリーポイント
│       └── main.go
├── internal/
│   └── api/             # APIハンドラー
│       ├── health.go    # ヘルスチェックAPI
│       └── files.go     # ファイル管理API
├── .devcontainer/       # Dev Container設定
│   └── devcontainer.json
├── go.mod
└── README.md
```

## 開発環境のセットアップ

### Dev Containerを使用する場合

1. Visual Studio Codeで本プロジェクトを開く
2. 拡張機能「Dev Containers」がインストールされていることを確認
3. コマンドパレット（Ctrl+Shift+P / Cmd+Shift+P）から「Dev Containers: Reopen in Container」を選択
4. コンテナが起動し、Go開発環境が自動的に構築されます

### ローカル環境で実行する場合

前提条件: Go 1.21以降がインストールされていること

```bash
# 依存関係のダウンロード
go mod download

# サーバーの起動
go run cmd/server/main.go
```

## APIエンドポイント

サーバーはデフォルトで `http://localhost:8080` で起動します。

### ヘルスチェック
- **URL**: `/api/health`
- **Method**: GET
- **レスポンス例**:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

### ファイル一覧取得
- **URL**: `/api/files`
- **Method**: GET
- **レスポンス例**:
```json
{
  "files": [
    {
      "name": "sample1.txt",
      "size": 1024,
      "path": "/files/sample1.txt"
    }
  ]
}
```

## テスト方法

```bash
# サーバー起動後、別のターミナルで以下を実行
curl http://localhost:8080/api/health
curl http://localhost:8080/api/files
```
