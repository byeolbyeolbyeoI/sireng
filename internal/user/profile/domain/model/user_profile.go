package model

// use join
type UserProfile struct {
	Id              int    `json:"id"`
	UserId          int    `json:"userId"`
	Username        string `json:"username"`
	ProfilePhotoUrl string `json:"profilePhotoUrl"`
	Role            int    `json:"role"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	CreatedAt       int    `json:"createdAt"`
	Bio             string `json:"bio"`
}

type UpdateUsernameRequest struct {
	Old string `json:"old" validate:"required,max=50"`
	New string `json:"new" validate:"required,max=50"`
}

type UpdateProfilePhotoRequest struct {
	Username           string `json:"username" validate:"required,max=50"`
	NewProfilePhotoURL string `json:"newProfilePhotoURL" validate:"required"`
}

type UpdateFirstNameRequest struct {
	Username     string `json:"username" validate:"required,max=50"`
	NewFirstName string `json:"newFirstName" validate:"required"`
}
type UpdateLastNameRequest struct {
	Username    string `json:"username" validate:"required,max=50"`
	NewLastName string `json:"newLastName" validate:"required"`
}

type UpdateBioRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	NewBio   string `json:"newBio" validate:"required"`
}
