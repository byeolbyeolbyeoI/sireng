package handler

import (
	"fmt"
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

type userHandlerImpl struct {
	service userService.UserService
	util    util.Util
}

func NewUserHandler(userService userService.UserService, util util.Util) UserHandler {
	return &userHandlerImpl{
		service: userService,
		util:    util,
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

	err = u.service.ValidateUserCredential(userCredential)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	/*
		// not good, because, if it doesn't exist, it will return fail to hash password
		// lowkey it's not that bad bcs the error is internal, the problem is the order will be messed up bcs why would u hash the password if the user doesn't exist? and u won't know that since it will returned first, buttttttt hashPasssowrd won't return error if there's no internal error so lowkey it is not that bad
		passwordChannel := make(chan string)
		go func() {
			hashedPassword, err := u.userService.HashPassword(userCredential.Password)
			if err != nil {
				u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
					Success: false,
					Message: err.Error(),
				})
				return
			}
			passwordChannel <- hashedPassword
		}()
	*/

	ok, err := u.service.IsExist(userCredential.Username)
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

	userCredential.Password, err = u.service.HashPassword(userCredential.Password)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// userCredential.Password = <-passwordChannel
	// input to db
	err = u.service.CreateUser(userCredential)
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

	err = u.service.ValidateUserCredential(userCredential)
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
	ok, err := u.service.IsExist(userCredential.Username)
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
	ok, err = u.service.IsCorrect(userCredential.Username, userCredential.Password)
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
	role, err := u.service.GetUserRole(userCredential.Username)
	if err != nil {
		u.util.WriteJSON(w, http.StatusUnauthorized, util.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	tokenString, err := u.service.GenerateTokenString(userCredential.Username, role)
	fmt.Println("Role : ", role)
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
