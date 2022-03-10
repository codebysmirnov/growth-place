package user

import (
	"github.com/google/uuid"

	"growth-place/application/domain"
)

// IUserRepo presents handlers methods
type IUserRepo interface {
	Create(user domain.User) error
	Update(user domain.User) error
	Read(id uuid.UUID) (domain.User, error)
}
