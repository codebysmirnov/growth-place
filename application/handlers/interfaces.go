package handlers

import (
	"github.com/google/uuid"

	"growth-place/application/services/user"
)

// IUserService presents user services methods
type IUserService interface {
	Create(login string, name, email, phone *string) error
	Authorization(email, password string) (user.AuthorizationView, error)
	PasswordEdit(id uuid.UUID, password string) error
}
