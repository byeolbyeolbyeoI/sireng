package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Input(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	// check
    fmt.Println("request.go:22", v)
	return nil
}
