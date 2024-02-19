package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenDuration  = 1 * time.Hour
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
