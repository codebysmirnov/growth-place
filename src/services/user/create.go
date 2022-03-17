package user

import (
	"github.com/dgrijalva/jwt-go"

	"growth-place/src/domain"
)

// CreateView user create response structure
type CreateView struct {
	Token     string `json:"token"`      // access token string
	ExpiredAt int64  `json:"expired_at"` // token deadline date time
}

// Create creates new user on system
func (s UserService) Create(login string, name, email, phone *string, password string) (CreateView, error) {
	logger := s.logger.With().Str("Method", "Create (User)").Logger()
	user, err := domain.NewUser(login, name, email, phone, password)
	if err != nil {
		logger.Error().Err(err).Msg("error on domain.NewUser")
		return CreateView{}, err
	}
	err = s.userRepo.Create(user)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.Create()")
		return CreateView{}, err
	}

	claims := domain.NewClaims(user.ID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		logger.Error().Err(err).Msg("error on token.SignedString)")
		return CreateView{}, ErrOnTokenSigning
	}

	return CreateView{
		Token:     tokenString,
		ExpiredAt: claims.ExpiresAt,
	}, nil
}
