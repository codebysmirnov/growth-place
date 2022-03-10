package user

import (
	"github.com/google/uuid"

	"growth-place/application/helpers"
)

// PasswordEdit edit user password
func (s UserService) PasswordEdit(id uuid.UUID, password string) error {
	logger := s.logger.With().Str("Method", "PasswordEdit").Logger()
	user, err := s.userRepo.Read(id)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Read()")
		return err
	}
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		logger.Error().Err(err).Msg("helpers.HashPassword()")
		return err
	}
	user.Password = &hashedPassword
	err = s.userRepo.Update(user)
	if err != nil {
		logger.Error().Err(err).Msg("s.userRepo.Update()")
		return err
	}
	return nil
}
