package repository

import (
	"database/sql"
	userModel "github.com/chaaaeeee/sireng/internal/user/auth/domain/model"
	"github.com/chaaaeeee/sireng/util"
)

type userRepositoryImpl struct {
	db   *sql.DB
	util util.Util
}

var ()

func NewUserRepository(db *sql.DB, util util.Util) UserRepository {
	return &userRepositoryImpl{
		db:   db,
		util: util,
	}
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
	_, err := u.db.Exec("INSERT INTO users(username, password_hashed, role) VALUES (?,?,?)", userCredential.Username, userCredential.Password, "user")
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepositoryImpl) GetUserRoleByUsername(username string) (string, error) {
	var role string
	err := u.db.QueryRow("SELECT role FROM users WHERE username=?", username).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
