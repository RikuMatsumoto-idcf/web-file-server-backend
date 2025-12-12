package api

import (
	"encoding/json"
	"net/http"
)

// FileInfo はファイル情報の構造体
type FileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

// FilesResponse はファイル一覧のレスポンス構造体
type FilesResponse struct {
	Files []FileInfo `json:"files"`
}

// FilesHandler はファイル一覧エンドポイント
func FilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// サンプルデータ
	response := FilesResponse{
		Files: []FileInfo{
			{Name: "sample1.txt", Size: 1024, Path: "/files/sample1.txt"},
			{Name: "sample2.txt", Size: 2048, Path: "/files/sample2.txt"},
		},
	}

	json.NewEncoder(w).Encode(response)
}
