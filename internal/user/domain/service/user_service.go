package service

type UserService interface {
	IsExist(string) (bool, error)
	IsCorrect(string, string) (bool, error)
}
