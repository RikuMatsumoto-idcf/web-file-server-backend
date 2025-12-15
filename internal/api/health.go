package api

import (
	"encoding/json"
	"net/http"
)

// HealthResponse はヘルスチェックのレスポンス構造体
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthHandler はヘルスチェックエンドポイント
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Message: "Server is running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
