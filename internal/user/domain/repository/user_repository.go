package repository

type UserRepository interface {
	IsExist(string) (bool, error)
	GetPasswordHashed(string) (string, error)
}
