package repository

import (
	"net/http"

	"growth-place/libs/liberror"
)

var (
	ErrUserWithPassedPhoneIsExists = &liberror.Error{
		Message:  "User with passed phone already exists",
		Code:     "USER_WITH_PASSED_PHONE_IS_EXISTS",
		HTTPCode: http.StatusConflict,
	}
	ErrUserWithPassedEmailIsExists = &liberror.Error{
		Message:  "User with passed email already exists",
		Code:     "USER_WITH_PASSED_EMAIL_IS_EXISTS",
		HTTPCode: http.StatusConflict,
	}
	ErrUserWithPassedLoginIsExists = &liberror.Error{
		Message:  "User with passed login already exists",
		Code:     "USER_WITH_PASSED_LOGIN_IS_EXISTS",
		HTTPCode: http.StatusBadRequest,
	}
	ErrUserNotFound = &liberror.Error{
		Message:  "User not found",
		Code:     "USER_NOT_FOUND",
		HTTPCode: http.StatusNotFound,
	}
)
