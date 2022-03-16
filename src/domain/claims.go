package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Claims user authorization data
type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}
