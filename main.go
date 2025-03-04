package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/sha1server/sha1Plus" // 使用模块路径导入
)

type HashRequest struct {
	Data string `json:"data"`
}

type HashResponse struct {
	Hash  string `json:"sha1"`
	Error string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/sha1", sha1Handler)
	http.HandleFunc("/health", healthCheck)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sha1Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req HashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	hasher := sha1Plus.NewSha1Plus()
	hasher.Update([]byte(req.Data))
	digest := hasher.Final()

	response := HashResponse{
		Hash: hex.EncodeToString(digest),
	}

	json.NewEncoder(w).Encode(response)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "OK")
}

func sendError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(HashResponse{Error: msg})
}
