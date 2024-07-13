package service

import (
	userRepository "github.com/chaaaeeee/sireng/internal/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepo userRepository.UserRepository
}

func NewUserService(userRepo userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
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
