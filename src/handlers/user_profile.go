package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
	"growth-place/middlewares"
	"growth-place/src/services/user"
)

// Profile user profile
// @Summary get profile
// @Description returns user personal data
// @Tags users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} ProfileResponse
// @Failure 400 {object} liberror.Error
// @Failure 404 {object} liberror.Error
// @Failure 500 {object} liberror.Error
// @Router /user [GET]
func (h UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := middlewares.MustUserID(ctx)
	res, err := h.userService.Profile(ctx, id)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	err = h.encodeProfileResponse(w, res)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	return
}

// encodeProfileResponse encode user profile response to ProfileResponse
func (h UserHandler) encodeProfileResponse(
	w http.ResponseWriter,
	res interface{},
) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(
		ProfileResponse{
			res.(user.ProfileView),
		},
	)
}

// ProfileResponse response profile struct
type ProfileResponse struct {
	user.ProfileView
}
