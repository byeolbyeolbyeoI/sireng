package service

import (
	"errors"
	trackerModel "github.com/chaaaeeee/sireng/internal/tracker/domain/model"
	trackerRepository "github.com/chaaaeeee/sireng/internal/tracker/domain/repository"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
)

type trackerServiceImpl struct {
	repo     trackerRepository.TrackerRepository
	util     util.Util
	validate *validator.Validate
}

func NewTrackerService(trackerRepo trackerRepository.TrackerRepository, util util.Util, validate *validator.Validate) TrackerService {
	return &trackerServiceImpl{
		repo:     trackerRepo,
		util:     util,
		validate: validate,
	}
}

func (t *trackerServiceImpl) ValidateParam(userId int) error {
	if err := t.validate.Var(userId, "required,min=0"); err != nil {
		// check if err is
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}

		return err.(validator.ValidationErrors)
	}

	return nil
}

func (t *trackerServiceImpl) GetStudySessions() ([]trackerModel.StudySession, error) {
	studySessions, err := t.repo.GetStudySessions()
	if err != nil {
		return nil, err
	}

	return studySessions, nil
}

func (t *trackerServiceImpl) GetStudySessionsByUserId(userId int) ([]trackerModel.StudySession, error) {
	studySessions, err := t.repo.GetStudySessionsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return studySessions, nil
}

func (t *trackerServiceImpl) IsSessionActiveByUserId(userId int) (bool, error) {
	isActive, err := t.repo.IsSessionActiveByUserId(userId)
	if err != nil {
		return false, err
	}

	if isActive {
		return true, nil
	}

	return false, nil
}

func (t *trackerServiceImpl) CreateStudySession(studySessionRequest trackerModel.StudySessionRequest) error {
	err := t.repo.CreateStudySession(studySessionRequest)
	if err != nil {
		return err
	}

	return nil
}

func (t *trackerServiceImpl) EndStudySession(userId int) error {
	err := t.repo.EndStudySession(userId)
	if err != nil {
		return err
	}

	return nil
}

func (t *trackerServiceImpl) ValidateStudySessionRequest(studySessionRequest trackerModel.StudySessionRequest) error {
	//validate
	err := t.validate.Struct(studySessionRequest)
	if err != nil {
		return err
	}
	return nil
}
