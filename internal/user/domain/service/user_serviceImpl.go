package service

import (
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userRepository "github.com/chaaaeeee/sireng/internal/user/domain/repository"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImpl struct {
	userRepo  userRepository.UserRepository
	util      util.Util
	validator *validator.Validate
}

func NewUserService(userRepo userRepository.UserRepository, util util.Util, validator *validator.Validate) UserService {
	return &userServiceImpl{
		userRepo:  userRepo,
		util:      util,
		validator: validator,
	}
}

func (u *userServiceImpl) ValidateCreateUserInput(userCredential userModel.UserCredential) error {
    //validate
	return nil
}

// i don't think this is effective, i think if the id is successfully returned we can put true here(?)
func (u *userServiceImpl) IsExist(username string) (bool, error) {
	ok, err := u.userRepo.IsExist(username)
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
	passwordHashed, err := u.userRepo.GetPasswordHashed(username)
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
	err := u.userRepo.InputUser(userCredential)
	if err != nil {
		return err
	}

	return nil
}

func CreateSession(username string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token
}

func SignToken(token *jwt.Token) (string, error) {
	tokenString, err := token.SignedString([]byte("tes"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
