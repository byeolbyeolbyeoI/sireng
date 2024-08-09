package repository

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
)

type TrackerRepository interface {
	GetStudySessionsByUserId(int) ([]trackerModel.StudySession, error)
	GetStudySessions() ([]trackerModel.StudySession, error)
	IsSessionActiveByUserId(int) (bool, error)
	CreateStudySession(trackerModel.StudySessionRequest) error
	EndStudySession(int) error
}
