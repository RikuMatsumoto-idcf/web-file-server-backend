-- ユーザーの初期データ投入
-- 注意: パスワードハッシュは開発用のプレースホルダーです。
-- 実際のbcryptハッシュは60文字で適切なソルトを含む必要があります。
-- 本番環境では bcrypt.GenerateFromPassword() などで生成した正しいハッシュを使用してください。
INSERT INTO users (username, email, password_hash) VALUES
    ('alice', 'alice@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'),
    ('bob', 'bob@example.com', '$2a$10$xGxRZIiTuWKjH3A8HG5etuRHE8hGGv6RYFxmTgmhGE6Yrf8oLKEKy'),
    ('charlie', 'charlie@example.com', '$2a$10$Wd3rKoNvA2sPKd2Q7sTabuSc5YGT8jAy5VnVWHQM3pUvHXQKDZVGa');

-- TODOの初期データ投入
-- 注意: 日付は例示用です。実際の開発では適宜更新してください。
INSERT INTO todos (user_id, title, description, completed, priority, due_date) VALUES
    (1, 'Go言語の学習', 'Web APIの基礎を学ぶ', false, 1, '2030-12-31'),
    (1, 'データベース設計', 'PostgreSQLのテーブル設計を行う', true, 2, '2030-06-30'),
    (1, 'REST API実装', 'CRUDエンドポイントを実装する', false, 1, '2030-01-15'),
    (2, 'プロジェクトドキュメント作成', 'READMEとAPI仕様書を書く', false, 0, '2030-01-10'),
    (2, 'テストコード作成', 'ユニットテストと統合テストを書く', false, 2, '2030-01-20'),
    (3, 'コードレビュー', 'チームメンバーのPRをレビューする', true, 1, '2030-06-15'),
    (3, 'デプロイ準備', 'CI/CDパイプラインの構築', false, 2, '2030-02-01');
