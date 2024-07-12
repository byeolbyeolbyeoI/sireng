package handler

import (
	"net/http"
)

type (
	TrackerHandlerImpl struct {
		// trackerService service.TrackerService
	}
)

/*
func NewTrackerHandlerImpl(service service.TrackerService) TrackerHandler {

	return &TrackerHandlerImpl{
		trackerService: service,
	}
}
*/

func (t *TrackerHandlerImpl) AddStudyTime(w http.ResponseWriter, r *http.Request) {

}
