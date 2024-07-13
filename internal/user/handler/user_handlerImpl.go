package handler

import (
	userModel "github.com/chaaaeeee/sireng/internal/user/domain/model"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

type userHandlerImpl struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) UserHandler {
	return &userHandlerImpl{userService: userService}
}

func (u *userHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var userCredential userModel.UserCredential
	err := util.Input(r, userCredential)
	if err != nil {
		http.Error(w, "Error retrieving input", http.StatusInternalServerError)
	}

	// does it exist though?
	ok, err := u.userService.IsExist(userCredential.Username)
	// check if there's error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// if there's no error, but ok is false then do this
	if !ok {
		http.Error(w, "User does not exist", http.StatusInternalServerError)
	}

	// if there's no error and ok is true then continue the code (cool)

	// well if it does, does it match up?? (same w above btw)
	ok, err = u.userService.IsCorrect(userCredential.Username, userCredential.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if !ok {
		http.Error(w, "Password Incorrect", http.StatusInternalServerError)
	}

	// ok u good my g
	util.WriteJSON(w, http.StatusOK, "User logged in successfully")
}
