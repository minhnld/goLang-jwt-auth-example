package models

import (
	"../../randomstrings"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Username, PasswordHash, Role string
}

// https://tools.ietf.org/html/rfc7519
type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

func GenerateCSRFSecret() (string, error) {
	return randomstrings.GenerateRandomString(32)
}
