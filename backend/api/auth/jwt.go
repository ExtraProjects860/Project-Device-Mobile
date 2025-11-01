package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func generateClaims(id uint, durationTime time.Duration) Claims {
	expirationTime := time.Now().Add(durationTime)
	return Claims{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "GoApp",
		},
	}
}

func GenerateAccessToken(id uint, jwtKey string, rememberMe bool) (string, error) {
	var claims Claims
	if rememberMe {
		claims = generateClaims(id, time.Duration(6)*time.Hour)
	} else {
		claims = generateClaims(id, time.Duration(1)*time.Hour)
	}

	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	accessToken, err := JWTtoken.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateAccessToken(tokenStr, jwtKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
