package handlers

import (
	"net/http"

	"growth-place/libs/liberror"
)

var (
	ErrUnmarshal = &liberror.Error{
		Message:  "Can't unmarshal request",
		Code:     "UNMARSHAL_ERROR",
		HTTPCode: http.StatusBadRequest,
	}
)
