package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const TOKEN_DEFAULT_DURATION = time.Minute * 10

type MyClaims struct {
	jwt.StandardClaims
	ID int
}

func GenerateToken(userId int, duration time.Duration) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		ID: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析 jwt token
func ParseToken(token string) (*MyClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("Secret"), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*MyClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
