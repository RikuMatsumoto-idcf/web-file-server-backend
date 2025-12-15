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

type errResponseWriter struct {
	base http.ResponseWriter
	err    error
}

func (w *errResponseWriter) Header() http.Header {
	return w.base.Header()
}

func (w *errResponseWriter) Write(_ []byte) (int, error) {
	return 0, w.err
}

func (w *errResponseWriter) WriteHeader(statusCode int) {
	w.base.WriteHeader(statusCode)
}

func TestHealthHandler_ReturnsOKJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Fatalf("expected Content-Type to contain application/json, got %q", contentType)
	}

	var body HealthResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if body.Status != "ok" {
		t.Fatalf("expected status %q, got %q", "ok", body.Status)
	}
	if body.Message != "Server is running" {
		t.Fatalf("expected message %q, got %q", "Server is running", body.Message)
	}
}

func TestHealthHandler_LogsWhenEncodeFails(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)

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
	HealthHandler(rw, req)

	if !strings.Contains(rw.Header().Get("Content-Type"), "application/json") {
		t.Fatalf("expected Content-Type to contain application/json, got %q", rw.Header().Get("Content-Type"))
	}

	logged := logBuf.String()
	if !strings.Contains(logged, "Failed to encode health response") {
		t.Fatalf("expected encode failure to be logged, got %q", logged)
	}
}
