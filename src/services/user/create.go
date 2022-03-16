package user

import (
	"growth-place/src/domain"
)

// Create creates new user on system
func (s UserService) Create(login string, name, email, phone *string) error {
	logger := s.logger.With().Str("Method", "Create (User)").Logger()
	user, err := domain.NewUser(login, name, email, phone, nil)
	if err != nil {
		logger.Error().Err(err).Msg("error on domain.NewUser")
		return err
	}
	err = s.userRepo.Create(user)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Create()")
		return err
	}

	return nil
}
