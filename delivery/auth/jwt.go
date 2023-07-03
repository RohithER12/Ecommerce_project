package auth

import (
	"70_Off/infrastructure/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email   string
	UserID  uint
	AdminID uint
	jwt.RegisteredClaims
}

func GenerateJWT(email string, userID uint, adminID uint) (string, error) {
	expireTime := time.Now().Add(60 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email:   email,
		UserID:  userID,
		AdminID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	})

	tokenString, err := token.SignedString([]byte(config.GetJWTConfig()))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTConfig()), nil
		},
	)
	if err != nil || !token.Valid {
		return claims, errors.New("not valid token")
	}
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return claims, errors.New("token expired re-login")
	}
	return claims, nil
}
