package user

import (
	"net/http"

	"growth-place/libs/liberror"
)

var (
	ErrOnTokenSigning = &liberror.Error{
		Message:  "Can't signing token",
		Code:     "SIGNING_TOKEN_ERROR",
		HTTPCode: http.StatusBadRequest,
	}
	ErrWrongPassword = &liberror.Error{
		Message:  "Wrong password",
		Code:     "PERMISSION_DENIED",
		HTTPCode: http.StatusUnauthorized,
	}
)
