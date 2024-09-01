package service

import (
	userProfileRepository "github.com/chaaaeeee/sireng/internal/user/profile/domain/repository"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
)

type userProfileServiceImpl struct {
	repo     userProfileRepository.UserProfileRepository
	util     util.Util
	validate *validator.Validate
}

func NewUserProfileService(userProfileRepo userProfileRepository.UserProfileRepository, util util.Util, validate *validator.Validate) UserProfileService {
	return &userProfileServiceImpl{
		repo:     userProfileRepo,
		util:     util,
		validate: validate,
	}
}

func (u *userProfileServiceImpl) ValidateRequest(r interface{}) error {
	err := u.validate.Struct(r)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileServiceImpl) UpdateUsername(old string, new string) error {
	err := u.repo.UpdateUsername(old, new)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileServiceImpl) UpdateProfilePhotoURL(username string, newProfilePhotoURL string) error {
	err := u.repo.UpdateProfilePhotoURL(username, newProfilePhotoURL)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileServiceImpl) UpdateFirstName(username string, newFirstName string) error {
	err := u.repo.UpdateFirstName(username, newFirstName)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileServiceImpl) UpdateLastName(username string, newLastName string) error {
	err := u.repo.UpdateLastName(username, newLastName)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileServiceImpl) UpdateBio(username string, newBio string) error {
	err := u.repo.UpdateBio(username, newBio)
	if err != nil {
		return err
	}

	return nil
}
