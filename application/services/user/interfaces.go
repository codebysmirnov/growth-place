package user

import "growth-place/application/domain"

// IUserRepo presents handlers methods
type IUserRepo interface {
	Create(user domain.User) error
}
