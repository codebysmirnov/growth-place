package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
)

// UserCreateArgs presents user create request arguments
type UserCreateArgs struct {
	Login string  `json:"login"` // login
	Name  *string `json:"name"`  // name
	Email *string `json:"email"` // email
	Phone *string `json:"phone"` // phone
}

// UserCreate handle user create request
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
