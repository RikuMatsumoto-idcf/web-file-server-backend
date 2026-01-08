package fs

// TODO:
// - Implement port.FileStore using local filesystem under a root directory.
// - Use atomic write (temp file + rename) if possible.
// - Translate OS errors to domain errors (e.g., not found).

