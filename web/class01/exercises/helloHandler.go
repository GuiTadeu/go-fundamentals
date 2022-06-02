package exercises

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HelloResponse{"Welcome, Guilherme!"}
	json.NewEncoder(w).Encode(response)
}

type HelloResponse struct {
	Message string
}