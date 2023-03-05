package middleware

import (
	"fmt"

	"mini-project-product/model/entity"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

func EncodeJwt(payload entity.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      payload.ID,
		"nama":    payload.Nama,
		"email":   payload.Email,
		"noTelp":  payload.NoTelp,
		"isAdmin": payload.IsAdmin,
	})

	hmacSecret := []byte("my_secret_key")
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJwt(tokenString string) (*entity.UserClaims, error) {

	if tokenString == "" {
		return nil, fmt.Errorf("Error token nil")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret := []byte("my_secret_key")
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := entity.UserClaims{}
		err := mapstructure.Decode(claims, &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, err
}

func IsAdmin(c *fiber.Ctx) (bool, error) {
	token := c.Get("token")
	decodedToken, err := DecodeJwt(token)
	fmt.Println("decodedToken", decodedToken)
	if err != nil {
		return false, err
	}

	return decodedToken.IsAdmin, nil
}
