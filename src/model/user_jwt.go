package model

import (
	"fmt"
	"os"
	"time"

	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET = "JWT_SECRET"
)

func (user *userDomain) GenerateToken() (string, *resterrors.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"id": user.id,
		"email": user.email,
		"name": user.name,
		"age": user.age,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", resterrors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token: %s", err.Error()),
		)
	}

	return tokenString, nil
}