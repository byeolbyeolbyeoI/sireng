package service

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerRepository "github.com/chaaaeeee/sireng/internal/tracker/domain/repository"
	"github.com/chaaaeeee/sireng/util"
)

type trackerServiceImpl struct {
	trackerRepo trackerRepository.TrackerRepository
	util        util.Util
}

func NewTrackerService(trackerRepo trackerRepository.TrackerRepository, util util.Util) TrackerService {
	return &trackerServiceImpl{
		trackerRepo: trackerRepo,
		util:        util,
	}
}

func (t *trackerServiceImpl) CreateSession(studySessionRequest trackerModel.StudySessionRequest) error {
	// check if active
	isActive, err := t.trackerRepo.IsSessionActiveByUserId(studySessionRequest.UserId)
	if err != nil {
		return err
	}

	if isActive {
		return ErrUserAlreadyInSession
	}

	err = t.trackerRepo.CreateSession(studySessionRequest)
	if err != nil {
		return err
	}

	return nil
}
