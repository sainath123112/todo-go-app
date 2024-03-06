package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(username string) (string, error) {
	secret := []byte("secretkey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iss":      time.Now(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString(secret)

	return signedToken, err
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, err
	} else {
		return true, nil
	}
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	return token, err
}
func GetUsername(tokenString string) string {
	token, err := ParseToken(tokenString)
	if err != nil {
		return err.Error()
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		fmt.Println("Error getting claims")
	}

	username := claims["username"].(string)
	return username
}
