package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
)

// UserCreateArgs presents user create request arguments
type UserCreateArgs struct {
	Login    string  `json:"login" example:"somelogin"`                                      // login
	Name     *string `json:"name" example:"some name"`                                       // name
	Email    *string `json:"email" example:"some@mail.com" minLength:"5" maxLength:"32"`     // email
	Phone    *string `json:"phone" example:"88009998889999"`                                 // phone
	Password string  `json:"password" example:"secret-word"  minLength:"7" maxLength:"1024"` // password
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

	err := h.userService.Create(args.Login, args.Name, args.Email, args.Phone, args.Password)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	err = h.encodeUserCreateResponse(w)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	return
}

// encodeUserCreateResponse encode user create response to NoContentResponse
func (h UserHandler) encodeUserCreateResponse(
	w http.ResponseWriter,
) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(NoContentResponse{})
}
