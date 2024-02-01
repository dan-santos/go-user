package model

import (
	"fmt"
	"os"
	"strings"
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

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}

func VerifyToken(tokenValue string) (UserDomainInterface, *resterrors.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, resterrors.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, resterrors.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, resterrors.NewUnauthorizedError("Invalid token")
	}

	return &userDomain{
		id: claims["id"].(string),
		email: claims["email"].(string),
		name: claims["name"].(string),
		age: int8(claims["age"].(float64)),
	}, nil
}