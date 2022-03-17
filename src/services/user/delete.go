package user

import (
	"context"

	"github.com/google/uuid"
)

// Delete mark user as deleted
func (s UserService) Delete(_ context.Context, id uuid.UUID) error {
	logger := s.logger.With().
		Str("Method", "Delete (User)").
		Str("UserID", id.String()).Logger()

	err := s.userRepo.Delete(id)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Delete()")
		return err
	}

	return nil
}
