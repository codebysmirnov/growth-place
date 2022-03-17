package valueobjects

import (
	"database/sql/driver"
	"net/http"
	"net/mail"
	"unicode/utf8"

	"growth-place/libs/liberror"
)

// Email presents email object type
type Email string

// Value implements sql.Valuer interface.
// Needed for automatic add to SQL-queries.
func (e Email) Value() (driver.Value, error) {
	return string(e), nil
}

var (
	ErrEmptyEmail = &liberror.Error{
		Message:  "Email can't be empty",
		Code:     "EMPTY_EMAIL",
		HTTPCode: http.StatusBadRequest,
	}
	ErrShortEmail = &liberror.Error{
		Message:  "Email must be more than 4 characters.",
		Code:     "SHORT_EMAIL",
		HTTPCode: http.StatusBadRequest,
	}
	ErrToLongEmail = &liberror.Error{
		Message:  "Email too long.",
		Code:     "LONG_EMAIL",
		HTTPCode: http.StatusBadRequest,
	}
)

// String get Email string
func (e Email) String() string {
	return string(e)
}

// Validate Email type
func (e Email) Validate() error {
	if e == "" {
		return ErrEmptyEmail
	}
	if len(e) < 5 {
		return ErrShortEmail
	}
	if utf8.RuneCountInString(e.String()) > 32 {
		return ErrToLongEmail
	}
	_, err := mail.ParseAddress(e.String())

	return err
}

// NewEmail returns new Email instance.
func NewEmail(email string) (Email, error) {
	e := Email(email)
	return e, e.Validate()
}
