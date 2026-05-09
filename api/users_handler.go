package api

import (
	"encoding/json"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"users": []string{}})
}
