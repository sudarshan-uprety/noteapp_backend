package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenDuration  = 1 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour
)

func GenerateTokens(userID int, email string) (string, string, time.Time, time.Time, error) {
	// Generate access
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["email"] = email
	accessTokenClaims["exp"] = time.Now().Add(accessTokenDuration).Unix()
	accessTokenClaims["iat"] = time.Now().Unix()

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	accessTokenExpTime := time.Unix(accessTokenClaims["exp"].(int64), 0)

	// Generate refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["user_id"] = userID
	refreshTokenClaims["email"] = email
	refreshTokenClaims["exp"] = time.Now().Add(refreshTokenDuration).Unix()
	refreshTokenClaims["iat"] = time.Now().Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}
	refreshTokenExpTime := time.Unix(refreshTokenClaims["exp"].(int64), 0)

	return accessTokenString, refreshTokenString, accessTokenExpTime, refreshTokenExpTime, nil
}

func ValidateAccessToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Extract and return the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func RefreshTokens(refreshTokenString string) (string, string, time.Time, time.Time, error) {
	// Parse the refresh token
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", "", time.Time{}, time.Time{}, fmt.Errorf("failed to parse refresh token: %v", err)
	}

	// Check if the refresh token is valid
	if !refreshToken.Valid {
		return "", "", time.Time{}, time.Time{}, fmt.Errorf("refresh token is not valid")
	}

	// Extract user details from the refresh token
	claims := refreshToken.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	email := claims["email"].(string)

	// Generate new tokens using the existing GenerateTokens function
	accessTokenString, refreshTokenString, accessTokenExpTime, refreshTokenExpTime, err := GenerateTokens(userID, email)
	if err != nil {
		return "", "", time.Time{}, time.Time{}, fmt.Errorf("failed to generate tokens: %v", err)
	}

	return accessTokenString, refreshTokenString, accessTokenExpTime, refreshTokenExpTime, nil
}
