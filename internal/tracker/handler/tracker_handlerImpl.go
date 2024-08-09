package handler

import (
	"errors"
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerService "github.com/chaaaeeee/sireng/internal/tracker/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
	"strconv"
)

type (
	TrackerHandlerImpl struct {
		service trackerService.TrackerService
		util    util.Util
	}
)

func NewTrackerHandler(service trackerService.TrackerService, util util.Util) TrackerHandler {
	return &TrackerHandlerImpl{
		service: service,
		util:    util,
	}
}

func (t *TrackerHandlerImpl) GetStudySessionsByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := t.service.ValidateParam(userId); err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	studySessions, err := t.service.GetStudySessionsByUserId(userId)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	t.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Study sessions retrieved successfully",
		Data:    studySessions,
	})
}

func (t *TrackerHandlerImpl) GetStudySessions(w http.ResponseWriter, r *http.Request) {
	studySessions, err := t.service.GetStudySessions()
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	t.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "Study sessions retrieved successfully",
		Data:    studySessions,
	})
}

func (t *TrackerHandlerImpl) CreateStudySession(w http.ResponseWriter, r *http.Request) {
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

	err = t.service.ValidateStudySessionRequest(studySessionRequest)
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	isActive, err := t.service.IsSessionActiveByUserId(studySessionRequest.UserId)
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
	err = t.service.CreateStudySession(studySessionRequest)
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

func (t *TrackerHandlerImpl) EndStudySession(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	// check for active session
	if err := t.service.ValidateParam(userId); err != nil {
		t.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	isActive, err := t.service.IsSessionActiveByUserId(userId)
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

	err = t.service.EndStudySession(userId)
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
