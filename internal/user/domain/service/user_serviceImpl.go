package service

import (
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userRepository "github.com/chaaaeeee/sireng/internal/user/domain/repository"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repo     userRepository.UserRepository
	util     util.Util
	validate *validator.Validate
}

func NewUserService(userRepo userRepository.UserRepository, util util.Util, validate *validator.Validate) UserService {
	return &userServiceImpl{
		repo:     userRepo,
		util:     util,
		validate: validate,
	}
}

func (u *userServiceImpl) ValidateUserCredential(userCredential userModel.UserCredential) error {
	//validate
	err := u.validate.Struct(userCredential)
	if err != nil {
		return err
	}
	return nil
}

// i don't think this is effective, i think if the id is successfully returned we can put true here(?)
func (u *userServiceImpl) IsExist(username string) (bool, error) {
	ok, err := u.repo.IsExist(username)
	// there's an error, or tons of errors
	if err != nil {
		return false, err
	}

	// no err but doesn't exist
	if !ok {
		return false, nil
	}

	// exists
	return true, nil
}

func (u *userServiceImpl) IsCorrect(username string, password string) (bool, error) {
	passwordHashed, err := u.repo.GetPasswordHashed(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userServiceImpl) HashPassword(password string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(passwordHashed), nil
}

func (u *userServiceImpl) CreateUser(userCredential userModel.UserCredential) error {
	err := u.repo.InputUser(userCredential)
	if err != nil {
		return err
	}

	return nil
}

func (u *userServiceImpl) GenerateTokenString(username string, role string) (string, error) {
	token := u.util.GenerateToken(username, role)

	tokenString, err := u.util.SignToken(token)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *userServiceImpl) GetUserRole(username string) (string, error) {
	role, err := u.repo.GetUserRoleByUsername(username)
	if err != nil {
		return "", err
	}

	return role, nil
}
