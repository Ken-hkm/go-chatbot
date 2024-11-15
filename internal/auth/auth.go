package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-chatbot/internal/db/models"
	"log"
	"os"
	"strconv"
	"time"
)

// ValidateToken validates the JWT token, returns the user ID, and checks for expiration.
func ValidateToken(tokenString string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	// Remove "Bearer " from the token if it's included in the header
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is what we expect (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	log.Println(token.Valid)

	// Verbose error handling
	if err != nil {
		if vErr, ok := err.(*jwt.ValidationError); ok {
			if vErr.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", errors.New("that's not even a token")
			} else if vErr.Errors&jwt.ValidationErrorExpired != 0 {
				return "", errors.New("token has expired")
			} else if vErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", errors.New("token not valid yet")
			} else {
				return "", errors.New("couldn't handle this token")
			}
		}
		return "", errors.New("invalid token")
	}

	// Extract claims from the token (assuming you put user_id and exp in the claims)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("user_id type: %T, value: %v", claims["user_id"], claims["user_id"])

		userID, ok := claims["user_id"].(string)
		if !ok {
			if userIDInt, ok := claims["user_id"].(float64); ok { // JWT claims are typically stored as float64
				userID = strconv.FormatFloat(userIDInt, 'f', -1, 64) // Convert to string
			} else {
				return "", errors.New("invalid user_id claim")
			}
		}
		// Check if the token is expired
		exp, ok := claims["exp"].(float64) // exp is typically stored as a float64
		if !ok {
			return "", errors.New("invalid exp claim")
		}

		// Compare token expiration with current time
		if int64(exp) < time.Now().Unix() {
			return "", errors.New("token has expired")
		}

		return userID, nil
	}

	return "", errors.New("invalid token claims")
}

func GenerateToken(user *models.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
