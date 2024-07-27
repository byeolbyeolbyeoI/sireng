package handler

import (
	"errors"
	"fmt"
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerService "github.com/chaaaeeee/sireng/internal/tracker/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
	"time"
)

type (
	TrackerHandlerImpl struct {
		trackerService trackerService.TrackerService
		util           util.Util
	}
)

func NewTrackerHandler(service trackerService.TrackerService, util util.Util) TrackerHandler {
	return &TrackerHandlerImpl{
		trackerService: service,
		util:           util,
	}
}

func (t *TrackerHandlerImpl) GetStudySessionsByUserIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (t *TrackerHandlerImpl) CreateStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	var studySessionRequest trackerModel.StudySessionRequest
	// take input
	err := t.util.Input(r, &studySessionRequest)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	//create session
	err = t.trackerService.CreateSession(studySessionRequest)
	if err != nil {
		// is in session
		if errors.Is(err, trackerService.ErrUserAlreadyInSession) {
			t.util.WriteJSON(w, http.StatusConflict, util.Response{
				Success: false,
				Message: trackerService.ErrUserAlreadyInSession.Error(),
			})
			return
		}
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	t.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Study Session Created Successfully",
	})
}
