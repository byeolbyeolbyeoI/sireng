package repository

type UserProfileRepository interface {
	UpdateUsername(string, string) error
	UpdateProfilePhotoURL(string, string) error
	UpdateFirstName(string, string) error
	UpdateLastName(string, string) error
	UpdateBio(string, string) error
}
