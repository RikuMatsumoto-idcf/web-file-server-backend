-- ユーザーの初期データ投入
INSERT INTO users (username, email, password_hash) VALUES
    ('alice', 'alice@example.com', '$2a$10$abcdefghijklmnopqrstuvwxyz123456789'), -- ダミーハッシュ
    ('bob', 'bob@example.com', '$2a$10$abcdefghijklmnopqrstuvwxyz987654321'),
    ('charlie', 'charlie@example.com', '$2a$10$abcdefghijklmnopqrstuvwxyz555555555');

-- TODOの初期データ投入
INSERT INTO todos (user_id, title, description, completed, priority, due_date) VALUES
    (1, 'Go言語の学習', 'Web APIの基礎を学ぶ', false, 1, '2024-12-31'),
    (1, 'データベース設計', 'PostgreSQLのテーブル設計を行う', true, 2, '2024-12-20'),
    (1, 'REST API実装', 'CRUDエンドポイントを実装する', false, 1, '2025-01-15'),
    (2, 'プロジェクトドキュメント作成', 'READMEとAPI仕様書を書く', false, 0, '2025-01-10'),
    (2, 'テストコード作成', 'ユニットテストと統合テストを書く', false, 2, '2025-01-20'),
    (3, 'コードレビュー', 'チームメンバーのPRをレビューする', true, 1, '2024-12-15'),
    (3, 'デプロイ準備', 'CI/CDパイプラインの構築', false, 2, '2025-02-01');
