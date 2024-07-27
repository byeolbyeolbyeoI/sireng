package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (u *utilImpl) WriteJSON(w http.ResponseWriter, code int, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	body, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), code)
	}
}
