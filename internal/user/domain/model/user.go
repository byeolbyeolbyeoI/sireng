package model

type User struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	PasswordHashed  string `json:"passwordHashed"`
	ProfilePhotoUrl string `json:"profilePhotoUrl"`
	RoleId          int    `json:"roleId"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	CreatedAt       int    `json:"createdAt"`
	Bio             string `json:"bio"`
}

type UserCredential struct {
	Username string `json:"username" validate:"required,max=20"`
	Password string `json:"password" validate:"required"`
}

type Role struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}
