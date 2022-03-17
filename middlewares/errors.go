package middlewares

import (
	"net/http"

	"growth-place/libs/liberror"
)

var (
	ErrUnauthorized = &liberror.Error{
		Message:  "Access denied.",
		Code:     "UNAUTHORIZED",
		HTTPCode: http.StatusUnauthorized,
	}
	ErrInvalidToken = &liberror.Error{
		Message:  "Token is invalid.",
		Code:     "INVALID_TOKEN",
		HTTPCode: http.StatusBadRequest,
	}
	ErrTokenProblem = &liberror.Error{
		Message:  "Can't parse token.",
		Code:     "INVALID_TOKEN",
		HTTPCode: http.StatusBadRequest,
	}
	ErrTokenScheme = &liberror.Error{
		Message:  "Token scheme is invalid",
		Code:     "INVALID_TOKEN_SCHEME",
		HTTPCode: http.StatusBadRequest,
	}
)
