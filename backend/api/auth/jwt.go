package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type RequestRefresh struct {
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func generateClaims(id uint, valueTime uint, timeFormat time.Duration) Claims {
	expirationTime := time.Now().Add(time.Duration(valueTime) * timeFormat)
	return Claims{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "GoApp",
		},
	}
}

func GenerateAccessToken(id uint, jwtKey string) (string, error) {
	claims := generateClaims(id, 5, time.Hour)

	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	accessToken, err := JWTtoken.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(id uint, refreshKey string) (string, error) {
	claims := generateClaims(id, 24, time.Hour)

	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	refreshToken, err := JWTtoken.SignedString([]byte(refreshKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func ValidateAccessToken(tokenStr, jwtKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
}

func ValidateRefreshToken(tokenStr string, refreshKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshKey), nil
	})
}

func RefreshToken(tokenStr, jwtKey, refreshKey string) (string, error) {
	token, err := ValidateRefreshToken(tokenStr, refreshKey)
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not parse claims")
	}

	newJWT, _, err := GenerateTokens(
		uint(claims["sub"].(float64)),
		jwtKey,
		refreshKey,
	)
	if err != nil {
		return "", fmt.Errorf("could not generate new token: %v", err)
	}

	return newJWT, nil
}

func GenerateTokens(id uint, jwtKey, refreshKey string) (string, string, error) {
	accessToken, err := GenerateAccessToken(id, jwtKey)
	if err != nil {
		return "", "", fmt.Errorf("could not create token: %v", err)
	}

	refreshToken, err := GenerateRefreshToken(id, refreshKey)
	if err != nil {
		return "", "", fmt.Errorf("could not create refresh token: %v", err)
	}

	return accessToken, refreshToken, nil
}
