package handler

import (
	"github.com/chaaaeeee/sireng/internal/user/profile/domain/model"
	userProfileService "github.com/chaaaeeee/sireng/internal/user/profile/domain/service"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

type userProfileHandlerImpl struct {
	service userProfileService.UserProfileService
	util    util.Util
}

func NewUserProfileHandler(userProfileService userProfileService.UserProfileService, util util.Util) UserProfileHandler {
	return &userProfileHandlerImpl{
		service: userProfileService,
		util:    util,
	}
}

func (u *userProfileHandlerImpl) UpdateUsername(w http.ResponseWriter, r *http.Request) {
	var updateUsernameRequest model.UpdateUsernameRequest

	err := u.util.Input(r, &updateUsernameRequest)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := u.service.ValidateRequest(updateUsernameRequest); err != nil {
		u.util.WriteJSON(w, http.StatusBadRequest, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.service.UpdateUsername(updateUsernameRequest.Old, updateUsernameRequest.New)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "aight",
	})
}

func (u *userProfileHandlerImpl) UpdateProfilePhoto(w http.ResponseWriter, r *http.Request) {
	var updateProfilePhotoRequest model.UpdateProfilePhotoRequest

	err := u.util.Input(r, &updateProfilePhotoRequest)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := u.service.ValidateRequest(updateProfilePhotoRequest); err != nil {
		u.util.WriteJSON(w, http.StatusBadRequest, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.service.UpdateProfilePhotoURL(updateProfilePhotoRequest.Username, updateProfilePhotoRequest.NewProfilePhotoURL)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "aight",
	})
}

func (u *userProfileHandlerImpl) UpdateFirstName(w http.ResponseWriter, r *http.Request) {
	var updateFirstNameRequest model.UpdateFirstNameRequest

	err := u.util.Input(r, &updateFirstNameRequest)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := u.service.ValidateRequest(updateFirstNameRequest); err != nil {
		u.util.WriteJSON(w, http.StatusBadRequest, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.service.UpdateFirstName(updateFirstNameRequest.Username, updateFirstNameRequest.NewFirstName)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "aight",
	})

}
func (u *userProfileHandlerImpl) UpdateLastName(w http.ResponseWriter, r *http.Request) {
	var updateLastNameRequest model.UpdateLastNameRequest

	err := u.util.Input(r, &updateLastNameRequest)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := u.service.ValidateRequest(updateLastNameRequest); err != nil {
		u.util.WriteJSON(w, http.StatusBadRequest, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.service.UpdateLastName(updateLastNameRequest.Username, updateLastNameRequest.NewLastName)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "aight",
	})
}

func (u *userProfileHandlerImpl) UpdateBio(w http.ResponseWriter, r *http.Request) {
	var updateBioRequest model.UpdateBioRequest

	err := u.util.Input(r, &updateBioRequest)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := u.service.ValidateRequest(updateBioRequest); err != nil {
		u.util.WriteJSON(w, http.StatusBadRequest, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = u.service.UpdateBio(updateBioRequest.Username, updateBioRequest.NewBio)
	if err != nil {
		u.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	u.util.WriteJSON(w, http.StatusOK, util.Response{
		Success: true,
		Message: "aight",
	})
}
