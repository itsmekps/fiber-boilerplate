package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// func ValidateToken(token)(token string) bool {

// }