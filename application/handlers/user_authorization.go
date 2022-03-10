package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
)

// UserAuthorizationArgs presents authorization arguments
type UserAuthorizationArgs struct {
	Email    string `json:"email"`    // email
	Password string `json:"password"` // password
}

// UserAuthorization handle user create request
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
