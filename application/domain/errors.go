package domain

import (
	"net/http"

	"growth-place/libs/liberror"
)

var (
	ErrUserEmptyLogin = &liberror.Error{
		Message:  "User login can't be empty",
		Code:     "USER_EMPTY_LOGIN",
		HTTPCode: http.StatusBadRequest,
	}
	ErrUserEmptyName = &liberror.Error{
		Message:  "User name can't be empty",
		Code:     "USER_EMPTY_NAME",
		HTTPCode: http.StatusBadRequest,
	}
	ErrUserEmptyEmail = &liberror.Error{
		Message:  "User email can't be empty",
		Code:     "USER_EMPTY_EMAIL",
		HTTPCode: http.StatusBadRequest,
	}
)
