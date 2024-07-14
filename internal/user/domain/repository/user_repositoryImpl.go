package repository

import (
	"database/sql"
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
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
	err := u.db.QueryRow("SELECT id FROM users WHERE username=?", username).Scan(&id)
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
	err := u.db.QueryRow("SELECT password_hashed FROM users WHERE username=?", username).Scan(&passwordHashed)
	if err != nil {
		return "", err
	}

	return passwordHashed, nil
}

func (u *userRepositoryImpl) InputUser(userCredential userModel.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO users(username, password_hashed, role_id) VALUES (?,?,?)", userCredential.Username, userCredential.Password, 1)
	if err != nil {
		return err
	}

	return nil
}
