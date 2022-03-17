package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
	"growth-place/src/services/user"
)

// UserAuthorizationArgs presents authorization arguments
type UserAuthorizationArgs struct {
	Login    string `json:"login" example:"some login"`       // login
	Password string `json:"password" example:"some_password"` // password
}

// UserAuthorization handle user authorization
// @Summary authorization
// @Description authorization user on system by email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param authorization body UserAuthorizationArgs true "User auth data"
// @Success 201 {object} UserAuthorizationResponse
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

	res, err := h.userService.Authorization(args.Login, args.Password)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	err = h.encodeUserAuthorizationResponse(w, res)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}
	return
}

// encodeUserAuthorizationResponse encode user authorization response to UserAuthorizationResponse
func (h UserHandler) encodeUserAuthorizationResponse(
	w http.ResponseWriter,
	res interface{},
) error {
	w.Header().Set("Content-Type", "src/json")
	return json.NewEncoder(w).Encode(
		UserAuthorizationResponse{
			res.(user.AuthorizationView),
		},
	)
}

// UserAuthorizationResponse presents user authorization response struct
type UserAuthorizationResponse struct {
	user.AuthorizationView
}
