package api

import (
	"encoding/json"
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
