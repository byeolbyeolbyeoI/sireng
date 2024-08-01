package handler

import (
	"net/http"
)

type TrackerHandler interface {
	GetStudySessionsByUserIdHandler(w http.ResponseWriter, r *http.Request)
	CreateStudySessionHandler(w http.ResponseWriter, r *http.Request)
	EndStudySessionHandler(w http.ResponseWriter, r *http.Request)
}
