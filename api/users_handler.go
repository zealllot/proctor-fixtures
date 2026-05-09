package api

import (
	"encoding/json"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(map[string]any{"users": []string{}})
}
