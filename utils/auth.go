package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

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

func VerifyToken(jwtToken string) error {
	parsedToken, err := jwt.Parse(jwtToken, parseTokenFn)

	// Verify is the parsed token is valid
	if err != nil {
		return errors.New("failed to parse the JWT Token")
	}
	if !parsedToken.Valid {
		return errors.New("failed to parse the JWT Token")
	}

	// Verify if the parsed token.Claims is the type of jwt.Claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("parsed claims are not type jwt.Claims")
	}

	// Server check of the information provided
	email := claims["email"]
	userId := claims["userId"]
	fmt.Println("Email: ", email, "User Id:", userId)
	return nil
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
