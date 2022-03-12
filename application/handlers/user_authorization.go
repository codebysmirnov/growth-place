package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
)

// UserAuthorizationArgs presents authorization arguments
type UserAuthorizationArgs struct {
	Email    string `json:"email" example:"some@mail.com"`    // email
	Password string `json:"password" example:"some_password"` // password
}

// UserAuthorization handle user authorization
// @Summary authorization
// @Description authorization user on system by email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param authorization body UserAuthorizationArgs true "User auth data"
// @Success 201 {object} user.AuthorizationView
// @Failure 400 {object} liberror.Error
// @Failure 404 {object} liberror.Error
// @Failure 500 {object} liberror.Error
// @Router /user/authorization [POST]
func (h UserHandler) UserAuthorization(w http.ResponseWriter, r *http.Request) {
	var args UserAuthorizationArgs
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		liberror.JSONError(w, ErrUnmarshal)
		return
	}

	res, err := h.userService.Authorization(args.Email, args.Password)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}
