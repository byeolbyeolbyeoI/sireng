package service

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
)

type TrackerService interface {
	IsSessionActiveByUserId(int) (bool, error)
	CreateStudySession(trackerModel.StudySessionRequest) error
	GetStudySessions() ([]trackerModel.StudySession, error)
	GetStudySessionsByUserId(int) ([]trackerModel.StudySession, error)
	EndStudySession(int) error
	ValidateStudySessionRequest(trackerModel.StudySessionRequest) error
	ValidateParam(int) error
}
