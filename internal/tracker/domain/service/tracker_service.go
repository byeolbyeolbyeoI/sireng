package service

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
)

type TrackerService interface {
	IsSessionActiveByUserId(int) (bool, error)
	CreateStudySession(trackerModel.StudySessionRequest) error
	EndStudySession(int) error
}
