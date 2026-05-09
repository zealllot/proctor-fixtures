package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	json.NewEncoder(w).Encode(map[string]any{
		"users": []string{},
		"limit": limit,
	})
}
