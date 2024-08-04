package handler

import (
	"fmt"
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

type userHandlerImpl struct {
	userService userService.UserService
	util        util.Util
}

func NewUserHandler(userService userService.UserService, util util.Util) UserHandler {
	return &userHandlerImpl{
		userService: userService,
		util:        util,
	}
}

// SignUp handles the user registration process.
// @Summary Register a new user
// @Description Registers a new user by taking user credentials, validating them, and storing them in the database.
// @Tags User
// @Accept json
// @Produce json
// @Param userCredential body userModel.UserCredential true "User Credentials"
// @Success 200 {object} util.Response "User signed up successfully"
// @Failure 409 {object} util.Response "Username already registered"
// @Failure 500 {object} util.Response "Internal Server Error"
// @Router /signup [post]

func (u *userHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	var userCredential userModel.UserCredential
	err := u.util.Input(r, &userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.userService.ValidateUserCredential(userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ok, err := u.userService.IsExist(userCredential.Username)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if ok {
		u.util.WriteJSON(w, http.StatusConflict, util.Response{
			Success: false,
			Message: "Username already registered",
		})
		return
	}

	// hash password
	userCredential.Password, err = u.userService.HashPassword(userCredential.Password)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// input to db
	err = u.userService.CreateUser(userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "User signed up successfully",
	})
}

func (u *userHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var userCredential userModel.UserCredential
	err := u.util.Input(r, &userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.userService.ValidateUserCredential(userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fmt.Println("user_handlerImpl.go:64", userCredential.Username, userCredential.Password)
	// does it exist though?
	// check if there's error
	ok, err := u.userService.IsExist(userCredential.Username)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// if there's no error, but ok is false then do this
	if !ok {
		u.util.WriteJSON(w, http.StatusNotFound, util.Response{
			Success: false,
			Message: "User does not exist",
		})
		return
	}

	// if there's no error and ok is true then continue the code (cool)

	// well if it does, does it match up?? (same w above btw)
	ok, err = u.userService.IsCorrect(userCredential.Username, userCredential.Password)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if !ok {
		u.util.WriteJSON(w, http.StatusUnauthorized, util.Response{
			Success: false,
			Message: "Password incorrect",
		})
		return
	}

	// ok u good my g
	// now jwt
	role, err := u.userService.GetUserRole(userCredential.Username)
	if err != nil {
		u.util.WriteJSON(w, http.StatusUnauthorized, util.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	tokenString, err := u.userService.GenerateTokenString(userCredential.Username, role)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "User logged in successfully",
		Data: map[string]string{
			"token": tokenString,
		},
	})
}
