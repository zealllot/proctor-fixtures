package api

import (
	"encoding/json"
	"net/http"
)

type ProfileReq struct {
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req ProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if len(req.DisplayName) > 100 {
		http.Error(w, "display name too long", 400)
		return
	}
	w.WriteHeader(204)
}
