package handler

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"answer": "User, hello!",
	}

	json.NewEncoder(w).Encode(response)
}
