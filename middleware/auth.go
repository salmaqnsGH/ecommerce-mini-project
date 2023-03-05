package middleware

import (
	"fmt"
	"os"

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

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	hmacSecret := []byte(JWTSecretKey)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJwt(tokenString string) (*entity.UserClaims, error) {

	isValid, token, err := CheckValidToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && isValid {
		user := entity.UserClaims{}
		err := mapstructure.Decode(claims, &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, err
}

func CheckValidToken(tokenString string) (bool, *jwt.Token, error) {
	if tokenString == "" {
		return false, nil, fmt.Errorf("Unauthorized")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
		hmacSecret := []byte(JWTSecretKey)
		return hmacSecret, nil
	})

	if err != nil {
		return false, nil, err
	}

	return token.Valid, token, nil
}

func GetUserData(c *fiber.Ctx) (*entity.UserClaims, error) {
	token := c.Get("token")
	decodedToken, err := DecodeJwt(token)
	if err != nil {
		return nil, err
	}

	return decodedToken, nil
}

func IsAdmin(c *fiber.Ctx) (bool, error) {
	userData, err := GetUserData(c)
	if err != nil {
		return false, err
	}

	return userData.IsAdmin, nil
}
