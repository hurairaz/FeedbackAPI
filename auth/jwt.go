package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type claims struct {
	UserID uint
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return "", err
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Println("JWT_SECRET_KEY is not set in the environment")
		return "", errors.New("JWT_SECRET_KEY is not set in the environment")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &claims{UserID: userID, RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    "hurairaz",
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
