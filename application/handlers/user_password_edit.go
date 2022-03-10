package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"growth-place/libs/liberror"
)

// UserPasswordEditArgs presents user password edit request arguments
type UserPasswordEditArgs struct {
	ID       uuid.UUID `json:"id"`       // user identifier
	Password string    `json:"password"` // new password
}

// UserPasswordEdit user password edit method
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
