package repository

import (
	"database/sql"
	"github.com/chaaaeeee/sireng/util"
)

type userProfileRepositoryImpl struct {
	db   *sql.DB
	util util.Util
}

func NewUserProfileRepository(db *sql.DB, util util.Util) UserProfileRepository {
	return &userProfileRepositoryImpl{
		db:   db,
		util: util,
	}
}

func (u *userProfileRepositoryImpl) UpdateUsername(old string, new string) error {
	_, err := u.db.Exec("UPDATE users SET username = ? WHERE username = ?", new, old)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileRepositoryImpl) UpdateProfilePhotoURL(username string, newProfilePhotoURL string) error {
	return nil
}

func (u *userProfileRepositoryImpl) UpdateFirstName(username string, newFirstName string) error {
	return nil
}

func (u *userProfileRepositoryImpl) UpdateLastName(username string, newLastName string) error {
	return nil
}

func (u *userProfileRepositoryImpl) UpdateBio(username string, newBio string) error {
	return nil
}
