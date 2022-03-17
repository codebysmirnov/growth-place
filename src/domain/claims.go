package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Claims user access data.
type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

// NewClaims returns new Claims instance.
func NewClaims(id uuid.UUID) Claims {
	return Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
}
