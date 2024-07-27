package service

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
)

type TrackerService interface {
	CreateSession(trackerModel.StudySessionRequest) error
}
