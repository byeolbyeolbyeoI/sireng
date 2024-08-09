package handler

import (
	"net/http"
)

type TrackerHandler interface {
	GetStudySessionsByUserId(w http.ResponseWriter, r *http.Request)
	GetStudySessions(w http.ResponseWriter, r *http.Request)
	CreateStudySession(w http.ResponseWriter, r *http.Request)
	EndStudySession(w http.ResponseWriter, r *http.Request)
}
