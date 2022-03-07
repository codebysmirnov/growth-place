package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User presents user data instance
type User struct {
	ID       uuid.UUID `json:"id" gorm:"id"`                       // identifier
	Login    string    `json:"login" gorm:"login"`                 // login
	Name     *string   `json:"name" gorm:"name"`                   // name
	Email    *string   `json:"email" gorm:"email"`                 // email
	Phone    *string   `json:"phone" gorm:"phone"`                 // phone
	Password *string   `json:"password,omitempty" gorm:"password"` // password

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
	if email != nil && *email == "" {
		return User{}, ErrUserEmptyEmail
	}
	return User{
		ID:       uuid.New(),
		Login:    login,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}, nil
}
