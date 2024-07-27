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
	token := u.util.CreateSession(userCredential.Username)

	tokenString, err := u.util.SignToken(token)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// decided to use statelessness
	/*
		cookie := &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * 3600 * time.Second),
			MaxAge:   3600 * 24 * 30,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	*/

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "User logged in successfully",
		Data: map[string]string{
			"token": tokenString,
		},
	})
}

/*
func (u *userHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	}

	http.SetCookie(w, cookie)

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "User logged out successfully",
	})
}
*/
