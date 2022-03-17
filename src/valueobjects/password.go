package valueobjects

import (
	"database/sql/driver"
	"net/http"
	"unicode/utf8"

	"growth-place/libs/liberror"
	"growth-place/src/helpers/hashpassword"
)

// Password presents Password object type
type Password string

// Value implements sql.Valuer interface.
// Needed for automatic add to SQL-queries.
func (p Password) Value() (driver.Value, error) {
	return string(p), nil
}

var (
	ErrEmptyPassword = &liberror.Error{
		Message:  "Password can't be empty",
		Code:     "EMPTY_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
	ErrShortPassword = &liberror.Error{
		Message:  "Password must be more than 7 characters.",
		Code:     "SHORT_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
	ErrToLongPassword = &liberror.Error{
		Message:  "Password too long.",
		Code:     "LONG_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
	ErrWrongPassword = &liberror.Error{
		Message:  "Incorrect password specified.",
		Code:     "WRONG_PASSWORD",
		HTTPCode: http.StatusBadRequest,
	}
)

// String get Password string
func (p Password) String() string {
	return string(p)
}

// Validate Password type
func validate(p string) error {
	if p == "" {
		return ErrEmptyPassword
	}
	if len(p) < 5 {
		return ErrShortPassword
	}
	if utf8.RuneCountInString(p) > 1024 {
		return ErrToLongPassword
	}

	return nil
}

// NewPassword returns new Password instance.
// When creating a password, the password is hashed.
func NewPassword(password string) (Password, error) {
	if err := validate(password); err != nil {
		return "", err
	}
	hashed, err := hashpassword.GenerateHashFromPassword(password)
	if err != nil {
		return "", err
	}
	p := Password(hashed)
	return p, nil
}

// Compare сompares the password to the given string.
// If the passed string is not identical to the password,
// the check is considered unsuccessful.
func (p Password) Compare(password string) error {
	ok, err := hashpassword.ComparePasswordAndHash(password, p.String())
	if !ok {
		return ErrWrongPassword
	}
	return err
}
