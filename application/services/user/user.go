package user

import (
	"growth-place/application/repository"
)

// UserService presents service for manage user data
type UserService struct {
	userRepo repository.UserRepo
}

// NewUserService returns new UserService instance
func NewUserService(userRepo repository.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}
