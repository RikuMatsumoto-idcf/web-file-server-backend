package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerMux_RoutesToHandlers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", HealthHandler)
	mux.HandleFunc("/api/files", FilesHandler)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	// health
	{
		res, err := http.Get(srv.URL + "/api/health")
		if err != nil {
			t.Fatalf("GET /api/health failed: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status %d for /api/health, got %d", http.StatusOK, res.StatusCode)
		}

		var body HealthResponse
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode /api/health JSON: %v", err)
		}
		if body.Status != "ok" {
			t.Fatalf("expected health.status %q, got %q", "ok", body.Status)
		}
	}

	// files
	{
		res, err := http.Get(srv.URL + "/api/files")
		if err != nil {
			t.Fatalf("GET /api/files failed: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status %d for /api/files, got %d", http.StatusOK, res.StatusCode)
		}

		var body FilesResponse
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode /api/files JSON: %v", err)
		}
		if len(body.Files) == 0 {
			t.Fatalf("expected at least 1 file")
		}
	}
}
