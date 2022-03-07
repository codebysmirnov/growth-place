package user

import (
	"growth-place/application/domain"
)

// Create creates new user on system
func (s UserService) Create(login string, name, email, phone *string) error {
	user, err := domain.NewUser(login, name, email, phone, nil)
	if err != nil {
		return err
	}
	return s.userRepo.Create(user)
}
