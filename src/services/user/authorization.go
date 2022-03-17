package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"growth-place/src/domain"
)

// AuthorizationView user auth response structure
type AuthorizationView struct {
	Token     string `json:"token"`      // access token string
	ExpiredAt int64  `json:"expired_at"` // token deadline date time
}

// Authorization authorize user and makes user token
func (s UserService) Authorization(email, password string) (AuthorizationView, error) {
	logger := s.logger.With().
		Str("Method", "Authorozation (User)").
		Str("email", email).
		Logger()

	user, err := s.userRepo.ReadByMail(email)
	if err != nil {
		logger.Error().Err(err).Msg("error on s.userRepo.ReadByMail()")
		return AuthorizationView{}, err
	}

	err = user.Password.Compare(password)
	if err != nil {
		logger.Error().Err(err).Msg("error on hashpassword.ComparePasswordAndHash()")
		return AuthorizationView{}, err
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &domain.Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		logger.Error().Err(err).Msg("error on token.SignedString)")
		return AuthorizationView{}, ErrOnTokenSigning
	}

	return AuthorizationView{
		Token:     tokenString,
		ExpiredAt: expirationTime,
	}, nil
}
