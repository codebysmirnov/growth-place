package user

import (
	"context"

	"github.com/google/uuid"

	"growth-place/src/domain/valueobjects"
)

// PasswordEdit edit user password
func (s UserService) PasswordEdit(
	_ context.Context,
	id uuid.UUID,
	password string,
) error {
	logger := s.logger.With().Str("Method", "PasswordEdit").Logger()
	user, err := s.userRepo.Read(id)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Read()")
		return err
	}
	pass, err := valueobjects.NewPassword(password)
	if err != nil {
		logger.Error().Err(err).Msg("helpers.HashPassword()")
		return err
	}
	user.Password = pass
	err = s.userRepo.Update(user)
	if err != nil {
		logger.Error().Err(err).Msg("s.userRepo.Update()")
		return err
	}
	return nil
}
