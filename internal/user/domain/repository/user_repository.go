package repository

import (
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
)

type UserRepository interface {
	IsExist(string) (bool, error)
	GetPasswordHashed(string) (string, error)
	InputUser(userModel.UserCredential) error
}
