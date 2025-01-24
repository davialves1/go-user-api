package utils

import (
	"errors"
	"os"
	"time"
	"user/server/models"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId string) (string, error) {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return "", errors.New("JWT secret key not available to generate JWT Token")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(jwtKey))
}

func VerifyToken(jwtToken string) (models.UserJWT, error) {
	var user models.UserJWT
	parsedToken, err := jwt.Parse(jwtToken, parseTokenFn)

	// Verify is the parsed token is valid
	if err != nil {
		return user, errors.New("failed to parse the JWT Token")
	}
	if !parsedToken.Valid {
		return user, errors.New("failed to parse the JWT Token")
	}

	// Verify if the parsed token.Claims is the type of jwt.Claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("parsed claims are not type jwt.Claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return user, errors.New("missing email from jwt claim")
	}

	id, ok := claims["userId"].(string)
	if !ok {
		return user, errors.New("missing id from jwt claim")
	}

	return models.UserJWT{Email: email, Id: id}, nil
}

func parseTokenFn(jwtToken *jwt.Token) (interface{}, error) {
	// Verify if is the correct signing method
	_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid signing method")
	}

	// Retrieve Jwt key
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return nil, errors.New("JWT secret key not available to verify JWT Token")
	}
	return []byte(jwtKey), nil
}
