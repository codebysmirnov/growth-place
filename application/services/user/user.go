package user

import (
	"github.com/rs/zerolog"

	"growth-place/application/repository"
)

// UserService presents service for manage user data
type UserService struct {
	userRepo repository.UserRepo
	logger   zerolog.Logger
}

// NewUserService returns new UserService instance
func NewUserService(userRepo repository.UserRepo, logger zerolog.Logger) UserService {
	return UserService{userRepo: userRepo, logger: logger}
}