package helper

import (
	"time"

	config "Main/configs"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(userID uint, userRole string) (string, error) {
	tokenJWT := config.GetTokenJWTConfig()
	secretKey := []byte(tokenJWT.JWT)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
