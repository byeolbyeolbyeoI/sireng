package service

import (
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
)

type UserService interface {
	IsExist(string) (bool, error)
	IsCorrect(string, string) (bool, error)
	HashPassword(string) (string, error)
	CreateUser(userModel.UserCredential) error
	ValidateUserCredential(userModel.UserCredential) error
	GenerateTokenString(string, string) (string, error)
	GetUserRole(string) (string, error)
}
