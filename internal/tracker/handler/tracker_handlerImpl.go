package handler

import (
	"errors"
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerService "github.com/chaaaeeee/sireng/internal/tracker/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
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

	isActive, err := t.trackerService.IsSessionActiveByUserId(studySessionRequest.UserId)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if isActive {
		t.util.WriteJSON(w, http.StatusNotFound, util.Response{
			Success: false,
			Message: "User is currently studying",
		})
		return
	}

	//create session
	err = t.trackerService.CreateStudySession(studySessionRequest)
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
		Message: "Study session created successfully",
	})
}

func (t *TrackerHandlerImpl) EndStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	type requestStruct struct {
		UserId int `json:"userId"`
	}

	var request requestStruct

	err := t.util.Input(r, &request)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// check for active session

	isActive, err := t.trackerService.IsSessionActiveByUserId(request.UserId)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if !isActive {
		t.util.WriteJSON(w, http.StatusNotFound, util.Response{
			Success: false,
			Message: "User is not studying",
		})
		return
	}
	err = t.trackerService.EndStudySession(request.UserId)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	t.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Study session ended successfully",
	})
}
