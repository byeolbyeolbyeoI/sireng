package handler

import "net/http"

type UserProfileHandler interface {
	UpdateUsername(w http.ResponseWriter, r *http.Request)
	UpdateProfilePhoto(w http.ResponseWriter, r *http.Request)
	UpdateFirstName(w http.ResponseWriter, r *http.Request)
	UpdateLastName(w http.ResponseWriter, r *http.Request)
	UpdateBio(w http.ResponseWriter, r *http.Request)
}
