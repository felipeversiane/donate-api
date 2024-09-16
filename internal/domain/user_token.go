package domain

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/felipeversiane/donate-api/internal/infra/config/rest"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	JWT_REFRESH_SECRET_KEY = "JWT_REFRESH_SECRET_KEY"
	JWT_SECRET_KEY         = "JWT_SECRET_KEY"
)

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func (ud *UserDomain) GenerateToken() (string, string, *rest.RestError) {
	accessToken, err := ud.GenerateAccessToken()
	if err != nil {
		return "", "", err
	}
	refreshToken, err := ud.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (ud *UserDomain) GenerateAccessToken() (string, *rest.RestError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID.String(),
		"email": ud.Email,
		"name":  ud.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest.NewInternalServerError(
			fmt.Sprintf("error trying to generate JWT token, err=%s", err.Error()))
	}

	return tokenString, nil
}

func (ud *UserDomain) GenerateRefreshToken() (string, *rest.RestError) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	refreshTokenClaims := jwt.MapClaims{
		"id":  ud.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", rest.NewInternalServerError(
			fmt.Sprintf("error trying to generate refresh token, err=%s", err.Error()))
	}

	return refreshTokenString, nil
}

func VerifyAccessToken(tokenValue string) (*UserDomain, *rest.RestError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, rest.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest.NewUnauthorizedRequestError("invalid token")
	}

	id, _ := uuid.Parse(claims["id"].(string))

	return &UserDomain{
		ID:    id,
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
	}, nil
}

func VerifyRefreshToken(tokenValue string) (*UserDomain, *rest.RestError) {
	secret := os.Getenv(JWT_REFRESH_SECRET_KEY)

	token, err := jwt.Parse(removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, rest.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest.NewUnauthorizedRequestError("invalid token")
	}

	id, _ := uuid.Parse(claims["id"].(string))

	return &UserDomain{
		ID: id,
	}, nil
}
