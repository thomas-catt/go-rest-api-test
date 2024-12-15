package utils

import (
	"fmt"
	"rest-api/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func jwtSecret() []byte {
	return []byte("SECRETSTRINGFORRRJWT")
}

func CreateJWT(user types.TokenUser) string {
	secret := jwtSecret()

	claim := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		fmt.Println("Failed to sign JWT:", err)
		return ""
	}

	return tokenString
}

func DecodeJWT(jwtToken string) (types.TokenUser, bool) {
	tokenUser := types.TokenUser{}
	token, err := jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret(), nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT:", err)
		return tokenUser, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if token.Valid && ok {
		tokenUser.Id = int64(claims["id"].(float64))
		tokenUser.Email = claims["email"].(string)
	}

	return tokenUser, tokenUser.Id != 0
}
