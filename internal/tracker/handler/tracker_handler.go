package handler

import (
	"net/http"
)

type TrackerHandler interface {
	AddStudyTime(w http.ResponseWriter, r *http.Request)
}
