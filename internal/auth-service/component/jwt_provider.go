package jwt

import (
	"common"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenProvider interface {
	GenerateToken(payload *TokenPayload, expiry int) (*Token, error)
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
	ID   int
	Role string
}

type Claims struct {
	jwt.RegisteredClaims
	ID   int
	Role string
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

func (j *jwtProvider) GenerateToken(payload *TokenPayload, expiry int) (*Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:   payload.ID,
		Role: payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(expiry))),

			ID: fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	aToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   aToken,
		Created: now,
		Expiry:  expiry,
	}, nil
}

type CustomClaims struct {
	ID   int
	Role string
}

func (j *jwtProvider) Validate(myToken string) (*Claims, error) {
	var uClaims *Claims
	resp, err := jwt.ParseWithClaims(myToken, uClaims, func(t *jwt.Token) (interface{}, error) { return []byte(j.secret), nil })
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, common.ErrJWTExpired
		} else if err == jwt.ErrSignatureInvalid {
			return nil, common.InvalidJWTToken
		}
		return nil, err
	}

	if !resp.Valid {
		return nil, common.InvalidJWTToken //Invalid token
	}

	claims := resp.Claims.(*Claims)
	return claims, nil

}
