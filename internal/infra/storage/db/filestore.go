package db

// TODO:
// - Implement port.FileStore using Postgres.
// - Store both file name and binary content (Postgres: BYTEA).
// - Keep SQL/driver details in this package (infra).
// - Translate DB errors into domain errors as needed.
//   - Unique violation (SQLSTATE 23505) -> domain.ErrAlreadyExists
//   - No rows -> domain.ErrNotFound
//
// Table sketch (example):
//   CREATE TABLE IF NOT EXISTS files (
//     name TEXT PRIMARY KEY,
//     data BYTEA NOT NULL,
//     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
//     updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
//   );