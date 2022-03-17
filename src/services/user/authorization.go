package user

import (
	"github.com/dgrijalva/jwt-go"

	"growth-place/src/domain"
)

// AuthorizationView user auth response structure
type AuthorizationView struct {
	Token     string `json:"token"`      // access token string
	ExpiredAt int64  `json:"expired_at"` // token deadline date time
}

// Authorization authorize user and makes user token
func (s UserService) Authorization(login, password string) (AuthorizationView, error) {
	logger := s.logger.With().
		Str("Method", "Authorization (User)").
		Str("login", login).
		Logger()

	user, err := s.userRepo.ReadByLogin(login)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.ReadByLogin()")
		return AuthorizationView{}, err
	}

	err = user.Password.Compare(password)
	if err != nil {
		logger.Error().Err(err).Msg("error on hashpassword.ComparePasswordAndHash()")
		return AuthorizationView{}, err
	}

	claims := domain.NewClaims(user.ID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		logger.Error().Err(err).Msg("error on token.SignedString)")
		return AuthorizationView{}, ErrOnTokenSigning
	}

	return AuthorizationView{
		Token:     tokenString,
		ExpiredAt: claims.ExpiresAt,
	}, nil
}
