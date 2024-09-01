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
	_, err := u.db.Exec("UPDATE user_profile b INNER JOIN users a ON b.user_id = a.id SET b.first_name = ? WHERE a.username = ?", newFirstName, username)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileRepositoryImpl) UpdateLastName(username string, newLastName string) error {
	_, err := u.db.Exec("UPDATE user_profile b INNER JOIN users a ON b.user_id = a.id SET b.last_name = ? WHERE a.username = ?", newLastName, username)
	if err != nil {
		return err
	}

	return nil
}

func (u *userProfileRepositoryImpl) UpdateBio(username string, newBio string) error {
	_, err := u.db.Exec("UPDATE user_profile b INNER JOIN users a ON b.user_id = a.id SET b.bio = ? WHERE a.username = ?", newBio, username)
	if err != nil {
		return err
	}

	return nil
}
