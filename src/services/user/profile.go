package user

import (
	"context"
	"time"

	"github.com/google/uuid"

	"growth-place/src/domain/valueobjects"
)

// ProfileView presents profile response struct
type ProfileView struct {
	Login     string             `json:"login" example:"some-login"`                // login
	Name      *string            `json:"name" example:"some-name"`                  // name
	Email     valueobjects.Email `json:"email" example:"some@mail.com"`             // email
	Phone     *string            `json:"phone" example:"88009998889988"`            // phone
	CreatedAt time.Time          `json:"created_at" example:"0001-01-01T00:00:00Z"` // user create datetime
}

// Profile returns user data
func (s UserService) Profile(
	_ context.Context,
	id uuid.UUID,
) (ProfileView, error) {
	logger := s.logger.With().
		Str("Method", "Profile (User)").
		Str("UserID", id.String()).Logger()

	user, err := s.userRepo.Read(id)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Read()")
		return ProfileView{}, err
	}

	return ProfileView{
		Login:     user.Login,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
	}, err
}
