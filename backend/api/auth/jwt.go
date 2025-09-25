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

func GenerateJwtToken(id uint) (string, error) {
	claims := generateClaims(id, 15, time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	jwtToken, err := token.SignedString(jwtKey())
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GenerateRefreshToken(id uint) (string, error) {
	claims := generateClaims(id, 24, time.Hour)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	refreshJWT, err := refreshToken.SignedString(refreshKey())
	if err != nil {
		return "", err
	}

	return refreshJWT, nil
}

// TODO daria pra ser uma função só que recebe o parametro da chave, mas preguiça e teria de ficar experto com os imports
func ValidateAccessToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey(), nil
    })
}

func ValidateRefreshToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return refreshKey(), nil
    })
}

func RefreshToken(tokenStr string) (string, error) {
	token, err := ValidateRefreshToken(tokenStr)
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not parse claims")
	}

	newJWT, _, err := GenerateTokens(uint(claims["sub"].(float64)))
	if err != nil {
		return "", fmt.Errorf("could not generate new token: %v", err)
	}
	
	return newJWT, nil
}

func GenerateTokens(id uint) (string, string, error) {
	token, err := GenerateJwtToken(id)
	if err != nil {
		return "", "", fmt.Errorf("could not create token: %v", err)
	}

	refresh, err := GenerateRefreshToken(id)
	if err != nil {
		return "", "", fmt.Errorf("could not create refresh token: %v", err)
	}

	return token, refresh, nil
}
