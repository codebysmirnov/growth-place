package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
)

// UserCreateArgs presents user create request arguments
type UserCreateArgs struct {
	Login string  `json:"login" example:"somelogin"`      // login
	Name  *string `json:"name" example:"somename"`        // name
	Email *string `json:"email" example:"some@mail.com"`  // email
	Phone *string `json:"phone" example:"88009998889999"` // phone
}

// UserCreate handle user create request
// @Summary  new user
// @Description create new user on system
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body UserCreateArgs true "User data"
// @Success 201 {object} NoContentResponse
// @Failure 400 {object} liberror.Error
// @Failure 409 {object} liberror.Error
// @Failure 500 {object} liberror.Error
// @Router /user [POST]
func (h UserHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	var args UserCreateArgs
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		liberror.JSONError(w, ErrUnmarshal)
		return
	}

	err := h.userService.Create(args.Login, args.Name, args.Email, args.Phone)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{}{})
	return
}
