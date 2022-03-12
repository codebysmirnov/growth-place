package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"growth-place/libs/liberror"
)

// UserPasswordEditArgs presents user password edit request arguments
type UserPasswordEditArgs struct {
	ID       uuid.UUID `json:"id" example:"8cef2e64-fe20-4259-8295-cb907f43cc0a"` // user identifier
	Password string    `json:"password" example:"some_password"`                  // new password
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

	err := h.userService.PasswordEdit(args.ID, args.Password)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{}{})
	return
}
