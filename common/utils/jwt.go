package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Username    string   `json:"username"`
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	SSO         int      `json:"sso"`
	jwt.RegisteredClaims
}
type JWTClaimsSSO struct {
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	jwt.RegisteredClaims
}

func ClaimJWT(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token, err
}

func DecodeJWT(tokenString string) (*JWTClaimsSSO, error) {
	claims := &JWTClaimsSSO{}
	_, _, err := jwt.NewParser().ParseUnverified(tokenString, claims)
	return claims, err
}

func GenerateJWT(claims *JWTClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
