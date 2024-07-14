package handler

import (
	"fmt"
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type userHandlerImpl struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) UserHandler {
	return &userHandlerImpl{userService: userService}
}

func (u *userHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	var userCredential userModel.UserCredential
	err := util.Input(r, &userCredential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok, err := u.userService.IsExist(userCredential.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ok {
		http.Error(w, "Username already registered", http.StatusConflict)
		return
	}

	// hash password
	userCredential.Password, err = u.userService.HashPassword(userCredential.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// input to db
	err = u.userService.CreateUser(userCredential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User signed up successfully",
	})
}

func (u *userHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var userCredential userModel.UserCredential
	err := util.Input(r, &userCredential)
	if err != nil {
		http.Error(w, "Error retrieving input", http.StatusInternalServerError)
		return
	}

	fmt.Println("user_handlerImpl.go:64", userCredential.Username, userCredential.Password)
	// does it exist though?
	// check if there's error
	ok, err := u.userService.IsExist(userCredential.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if there's no error, but ok is false then do this
	if !ok {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	// if there's no error and ok is true then continue the code (cool)

	// well if it does, does it match up?? (same w above btw)
	ok, err = u.userService.IsCorrect(userCredential.Username, userCredential.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Password Incorrect", http.StatusUnauthorized)
		return
	}

	// ok u good my g
	util.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User logged in successfully",
	})
}
