package db

// TODO:
// - Add tests for the Postgres-backed FileStore.
// - Recommended approach:
//   - Integration tests that require a running Postgres (e.g., via docker compose).
//   - Or guard tests behind an env var like TEST_DB_DSN and skip if not set.
// - Verify Put/Get/Exists and error mapping (not found, already exists).