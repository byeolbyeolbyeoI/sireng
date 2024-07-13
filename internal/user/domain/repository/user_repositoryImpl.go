package repository

import (
	"database/sql"
)

type userRepositoryImpl struct {
	db *sql.DB
}

var ()

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// refactor ke service, repo hanya untuk deal w database
func (u *userRepositoryImpl) IsExist(username string) (bool, error) {
	var id int
	err := u.db.QueryRow("SELECT id FROM users WHERE username=?", username).Scan(id)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepositoryImpl) GetPasswordHashed(username string) (string, error) {
	var passwordHashed string
	err := u.db.QueryRow("SELECT password_hashed FROM users WHERE username=?", username).Scan(passwordHashed)
	if err != nil {
		return "", err
	}

	return passwordHashed, nil
}
