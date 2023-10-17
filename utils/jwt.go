package utils

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your-secret-key")
var activeTokens = make(map[string]string) // Map to store active tokens (key: email, value: token)

// GenerateJWT generates a new JWT token for the given email and role.
func GenerateJWT(email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates the JWT token and returns the user's role.
func ValidateJWT(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")

	// Extract email from the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)

		// Check if the token is active (exists in activeTokens map)
		storedToken, found := activeTokens[email]
		if !found || storedToken != tokenString {
			return "", errors.New("Invalid token or user has logged out")
		}

		role := claims["role"].(string)
		return role, nil
	}

	return "", errors.New("Invalid token")
}
