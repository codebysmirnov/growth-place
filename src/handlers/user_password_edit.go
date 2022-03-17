package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
	"growth-place/middlewares"
)

// UserPasswordEditArgs presents user password edit request arguments
type UserPasswordEditArgs struct {
	Password string `json:"password" example:"some_password"` // new password
}

// UserPasswordEdit user password edit method
// @Summary edit password
// @Description edit user password: add new or replace old password
// @Tags users
// @Accept  json
// @Produce  json
// @Param password body UserPasswordEditArgs true "New password"
// @Success 201 {object} NoContentResponse
// @Failure 400 {object} liberror.Error
// @Failure 404 {object} liberror.Error
// @Failure 500 {object} liberror.Error
// @Router /user/password [POST]
func (h UserHandler) UserPasswordEdit(w http.ResponseWriter, r *http.Request) {
	var args UserPasswordEditArgs
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		liberror.JSONError(w, ErrUnmarshal)
		return
	}
	ctx := r.Context()
	token := middlewares.MustUserID(ctx)
	err := h.userService.PasswordEdit(ctx, token, args.Password)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	err = h.encodeUserPasswordEditResponse(w)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	return
}

// encodeUserPasswordEditResponse encode user password edit response to NoContentResponse
func (h UserHandler) encodeUserPasswordEditResponse(
	w http.ResponseWriter,
) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(NoContentResponse{})
}
