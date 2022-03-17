package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"growth-place/libs/liberror"
	"growth-place/src/domain"
)

// UserID context key to user id
type UserID struct{}

// MustUserID return user id from context
func MustUserID(ctx context.Context) uuid.UUID {
	return ctx.Value(UserID{}).(uuid.UUID)
}

const bearerScheme = "Bearer"

// getToken get token from passed header
func getToken(header string) (string, error) {
	token := strings.Split(header, " ")
	if token[0] != bearerScheme {
		return "", ErrTokenScheme
	}
	if len(token) < 2 {
		return "", ErrUnauthorized
	}
	return token[1], nil
}

// NewAuthorization return user auth middleware
func NewAuthorization(jwtKey string, logger zerolog.Logger) mux.MiddlewareFunc {
	logger = logger.With().Str("Method", "NewAuthorization (Middleware)").Logger()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := getToken(r.Header.Get("Authorization"))
			if err != nil {
				logger.Error().Err(err).Msg("error on getToken()")
				liberror.JSONError(w, err)
				return
			}

			t, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtKey), nil
			})
			if err != nil {
				logger.Error().Err(err).Msg("can't parse token")
				liberror.JSONError(w, ErrTokenProblem)
				return
			}

			claims, ok := t.Claims.(*domain.Claims)
			if !t.Valid || !ok {
				logger.Error().Err(err).Msg("token is not valid")
				liberror.JSONError(w, ErrInvalidToken)
				return
			}

			ctx := context.WithValue(r.Context(), UserID{}, claims.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
