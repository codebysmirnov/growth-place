package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"growth-place/src/valueobjects"
)

// User presents user data instance
type User struct {
	ID       uuid.UUID          `json:"id" gorm:"id"`                       // identifier
	Login    string             `json:"login" gorm:"login"`                 // login
	Name     *string            `json:"name" gorm:"name"`                   // name
	Email    valueobjects.Email `json:"email" gorm:"email"`                 // email
	Phone    *string            `json:"phone" gorm:"phone"`                 // phone
	Password *string            `json:"password,omitempty" gorm:"password"` // password

	CreatedAt time.Time      `json:"created_at" gorm:"created_at"` // datetime of user create on system
	UpdatedAt time.Time      `json:"updated_at" gorm:"created_at"` // datetime of user data modify
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"` // datetime of user data delete from system
}

// NewUser return new User instance
func NewUser(
	login string,
	name *string,
	email *string,
	phone *string,
	password *string,
) (User, error) {
	login = strings.TrimSpace(login)
	if login == "" {
		return User{}, ErrUserEmptyLogin
	}
	if name != nil && *name == "" {
		return User{}, ErrUserEmptyName
	}

	var (
		e   valueobjects.Email
		err error
	)
	if email != nil {
		e, err = valueobjects.NewEmail(*email)
		if err != nil {
			return User{}, err
		}
	}
	return User{
		ID:       uuid.New(),
		Login:    login,
		Name:     name,
		Email:    e,
		Phone:    phone,
		Password: password,
	}, nil
}
