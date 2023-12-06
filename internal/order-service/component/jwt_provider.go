package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenProvider interface {
	Validate(myToken string) (*Claims, error)
	NewPayLoad(id int, role string) *TokenPayload
}
type jwtProvider struct {
	secret string
}

func NewJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type TokenPayload struct {
	ID   int    `json:"user_id"`
	Role string `json:"role"`
}

type Claims struct {
	jwt.RegisteredClaims
	ID   int    `json:"user_id"`
	Role string `json:"role"`
}
type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"token_created"`
	Expiry  int       `json:"token_expiry"`
}

func (j *jwtProvider) NewPayLoad(id int, role string) *TokenPayload {
	return &TokenPayload{
		ID:   id,
		Role: role,
	}
}

func (j *jwtProvider) Validate(myToken string) (*Claims, error) {
	// var claims *Claims
	token, err := jwt.ParseWithClaims(myToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err

}
