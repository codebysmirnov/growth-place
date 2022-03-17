package user

import (
	"github.com/rs/zerolog"

	"growth-place/src/repository"
)

// UserService presents service for manage user data
type UserService struct {
	userRepo IUserRepo
	logger   zerolog.Logger
	jwtKey   string
}

// NewUserService returns new UserService instance
func NewUserService(
	userRepo repository.UserRepo,
	logger zerolog.Logger,
	jwtKey string,
) UserService {
	return UserService{userRepo: userRepo, logger: logger, jwtKey: jwtKey}
}
