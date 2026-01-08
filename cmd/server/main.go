package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RikuMatsumoto-idcf/web-file-server-backend/internal/api"
)

func main() {
	// ルーターの設定
	mux := http.NewServeMux()

	// APIハンドラーの登録
	mux.HandleFunc("/api/files", api.FilesHandler)

	// サーバーの起動
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
