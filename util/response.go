package util

import (
	"encoding/json"
	"net/http"
)

func (u *utilImpl) WriteJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	body, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	if _, err = w.Write(body); err != nil {
		http.Error(w, err.Error(), code)
	}
}
