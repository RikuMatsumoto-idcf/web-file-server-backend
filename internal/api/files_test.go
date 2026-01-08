package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFilesHandler_ReturnsFilesJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/files", nil)
	rec := httptest.NewRecorder()

	FilesHandler(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Fatalf("expected Content-Type to contain application/json, got %q", contentType)
	}

	var body FilesResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if len(body.Files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(body.Files))
	}

	if body.Files[0].Name != "sample1.txt" {
		t.Fatalf("expected first file name %q, got %q", "sample1.txt", body.Files[0].Name)
	}
	if body.Files[0].Size != 1024 {
		t.Fatalf("expected first file size %d, got %d", int64(1024), body.Files[0].Size)
	}
	if body.Files[0].Path != "/files/sample1.txt" {
		t.Fatalf("expected first file path %q, got %q", "/files/sample1.txt", body.Files[0].Path)
	}
}

func TestFilesHandler_LogsWhenEncodeFails(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/files", nil)

	var logBuf bytes.Buffer
	prevOut := log.Writer()
	prevFlags := log.Flags()
	log.SetOutput(&logBuf)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(prevOut)
		log.SetFlags(prevFlags)
	})

	base := httptest.NewRecorder()
	rw := &errResponseWriter{base: base, err: errors.New("write failed")}
	FilesHandler(rw, req)

	if !strings.Contains(rw.Header().Get("Content-Type"), "application/json") {
		t.Fatalf("expected Content-Type to contain application/json, got %q", rw.Header().Get("Content-Type"))
	}

	logged := logBuf.String()
	if !strings.Contains(logged, "Failed to encode files response") {
		t.Fatalf("expected encode failure to be logged, got %q", logged)
	}
}

type errResponseWriter struct {
	base http.ResponseWriter
	err  error
}

func (w *errResponseWriter) Header() http.Header {
	return w.base.Header()
}

func (w *errResponseWriter) Write(p []byte) (int, error) {
	return 0, w.err
}

func (w *errResponseWriter) WriteHeader(statusCode int) {
	w.base.WriteHeader(statusCode)
}
