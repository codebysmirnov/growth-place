package handlers

import (
	"encoding/json"
	"net/http"

	"growth-place/libs/liberror"
	"growth-place/middlewares"
)

// Delete user delete from system (mark as deleted)
// @Summary delete user
// @Description mark user as deleted
// @Tags users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} NoContentResponse
// @Failure 400 {object} liberror.Error
// @Failure 404 {object} liberror.Error
// @Failure 500 {object} liberror.Error
// @Router /user [DELETE]
func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := middlewares.MustUserID(ctx)
	err := h.userService.Delete(ctx, id)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	err = h.encodeDeleteResponse(w)
	if err != nil {
		liberror.JSONError(w, err)
		return
	}

	return
}

// encodeDeleteResponse encode user delete response to NoContentResponse
func (h UserHandler) encodeDeleteResponse(
	w http.ResponseWriter,
) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(NoContentResponse{})
}
