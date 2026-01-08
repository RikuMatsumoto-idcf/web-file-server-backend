-- filesテーブルの初期データ投入
-- 注意: このプロジェクトは「アップロードされたファイルをDBに保存」するのが目的なので、
-- 初期データは最小限のサンプル1件だけ入れておきます（不要なら削除OK）。
-- "hello" (0x68656c6c6f) を hello.txt として保存
INSERT INTO files (name, data)
VALUES ('hello.txt', decode('68656c6c6f', 'hex'))
ON CONFLICT (name) DO NOTHING;
