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
│   ├── devcontainer.json
│   ├── docker-compose.yml
│   └── initdb/          # データベース初期化SQL
│       ├── 01_schema.sql
│       └── 02_seeds.sql
├── go.mod
└── README.md
```

## 開発環境のセットアップ

### Dev Containerを使用する場合（推奨）

このプロジェクトは、GoアプリケーションとPostgreSQLデータベースを含むDocker Compose環境で動作します。

1. Visual Studio Codeで本プロジェクトを開く
2. 拡張機能「Dev Containers」がインストールされていることを確認
3. コマンドパレット（Ctrl+Shift+P / Cmd+Shift+P）から「Dev Containers: Reopen in Container」を選択
4. コンテナが起動し、Go開発環境とPostgreSQLデータベースが自動的に構築されます

#### データベースの確認

コンテナ起動後、以下の方法でデータベースが正常に起動していることを確認できます：

**方法1: VSCode SQLToolsを使用**
1. VSCodeのサイドバーからSQLToolsアイコンをクリック
2. 新しい接続を作成（接続情報は下記の環境変数を参照）
3. テーブルとデータを確認

**方法2: コマンドラインから接続**
```bash
# コンテナ内のターミナルで実行
psql -h db -U user -d db

# テーブル一覧を表示
\dt

# usersテーブルのデータを確認
SELECT * FROM users;

# todosテーブルのデータを確認
SELECT * FROM todos;
```

**方法3: Docker Composeコマンドで確認**
```bash
# コンテナの状態を確認（ホストマシンから）
docker compose -f .devcontainer/docker-compose.yml ps

# dbコンテナのログを確認
docker compose -f .devcontainer/docker-compose.yml logs db
```

### データベース接続情報

アプリケーションからデータベースに接続する際は、以下の環境変数が利用可能です：

| 環境変数 | 値 |
|---------|-----|
| `DB_HOST` | `db` |
| `DB_PORT` | `5432` |
| `DB_USER` | `user` |
| `DB_PASSWORD` | `password` |
| `DB_NAME` | `db` |

**⚠️ セキュリティに関する注意**
- 上記の認証情報は開発環境用の簡易的なものです
- 本番環境では強力なパスワードと適切なセキュリティ設定を使用してください
- 本番環境では環境変数を安全に管理する仕組み（AWS Secrets Manager、Kubernetes Secretsなど）を使用してください

### ローカル環境で実行する場合

前提条件: Go 1.22以降がインストールされていること

```bash
# 依存関係のダウンロード
go mod download

# サーバーの起動
go run cmd/server/main.go
```

**注意**: ローカル環境で実行する場合、PostgreSQLデータベースは含まれません。

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

## データベーススキーマ

### usersテーブル
ユーザー情報を管理するテーブルです。

| カラム | 型 | 説明 |
|--------|-----|------|
| id | SERIAL | 主キー |
| username | VARCHAR(50) | ユーザー名（ユニーク） |
| email | VARCHAR(100) | メールアドレス（ユニーク） |
| password_hash | VARCHAR(255) | パスワードハッシュ |
| created_at | TIMESTAMP | 作成日時 |
| updated_at | TIMESTAMP | 更新日時 |

### todosテーブル
TODOタスクを管理するテーブルです。

| カラム | 型 | 説明 |
|--------|-----|------|
| id | SERIAL | 主キー |
| user_id | INTEGER | ユーザーID（外部キー） |
| title | VARCHAR(200) | タスクタイトル |
| description | TEXT | タスク詳細 |
| completed | BOOLEAN | 完了フラグ |
| priority | INTEGER | 優先度 |
| due_date | DATE | 期限日 |
| created_at | TIMESTAMP | 作成日時 |
| updated_at | TIMESTAMP | 更新日時 |
