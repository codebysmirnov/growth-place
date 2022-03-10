package handlers

import "github.com/google/uuid"

// IUserService presents user services methods
type IUserService interface {
	Create(login string, name, email, phone *string) error
	PasswordEdit(id uuid.UUID, password string) error
}
