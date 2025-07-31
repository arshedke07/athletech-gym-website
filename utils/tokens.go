package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/arshedke07/athletech/model"
)

func GenerateToken(userId int, userName string, role string) (string, error) {
	// init the claims struct for the jwt payload
	claims := model.Claims{
		UserId:   userId,
		Role:     role,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// get the token string after signing it with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
