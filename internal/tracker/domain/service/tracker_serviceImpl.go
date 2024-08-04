package service

import (
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerRepository "github.com/chaaaeeee/sireng/internal/tracker/domain/repository"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
)

type trackerServiceImpl struct {
	trackerRepo trackerRepository.TrackerRepository
	util        util.Util
	validate    *validator.Validate
}

func NewTrackerService(trackerRepo trackerRepository.TrackerRepository, util util.Util, validate *validator.Validate) TrackerService {
	return &trackerServiceImpl{
		trackerRepo: trackerRepo,
		util:        util,
		validate:    validate,
	}
}

func (t *trackerServiceImpl) IsSessionActiveByUserId(userId int) (bool, error) {
	isActive, err := t.trackerRepo.IsSessionActiveByUserId(userId)
	if err != nil {
		return false, err
	}

	if isActive {
		return true, nil
	}

	return false, nil
}

func (t *trackerServiceImpl) CreateStudySession(studySessionRequest trackerModel.StudySessionRequest) error {
	err := t.trackerRepo.CreateStudySession(studySessionRequest)
	if err != nil {
		return err
	}

	return nil
}

func (t *trackerServiceImpl) EndStudySession(userId int) error {
	err := t.trackerRepo.EndStudySession(userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *trackerServiceImpl) ValidateStudySessionRequest(studySessionRequest trackerModel.StudySessionRequest) error {
	//validate
	err := u.validate.Struct(studySessionRequest)
	if err != nil {
		return err
	}
	return nil
}
