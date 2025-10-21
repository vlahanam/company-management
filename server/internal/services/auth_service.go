package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

const (
	expAssetToken   = 15 * time.Minute
	expRefreshToken = 7 * 24 * time.Hour
)

type authService struct {
	es            *userService
	accessSecret  string
	refreshSecret string
}

func NewAuthService(es *userService, accessSecret, refreshSecret string) *authService {
	return &authService{
		es:            es,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (as *authService) Login(ctx context.Context, data *requests.LoginRequest) (*models.Auth, error) {
	e, err := as.es.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, common.ErrorValidation.Clone().SetDetail("email", models.ErrEmailNotFound.Error())
	}

	checkPassword := common.CheckPasswordHash(data.Password, e.HashPassword)
	if !checkPassword {
		return nil, common.ErrorValidation.Clone().SetDetail("password", models.ErrInvalidPassword.Error())
	}

	auth, err := as.GenerateTokens(e.Email)
	if err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapError(err)
	}

	return auth, nil
}

func (as *authService) GenerateTokens(email string) (*models.Auth, error) {
	// Access Token (15 minutes)
	accessClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expAssetToken).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err := accessToken.SignedString([]byte(as.accessSecret))
	if err != nil {
		return nil, err
	}

	// Refresh Token (7 days)
	refreshClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expRefreshToken).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err := refreshToken.SignedString([]byte(as.refreshSecret))
	if err != nil {
		return nil, err
	}

	auth := &models.Auth{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return auth, nil
}

func (as *authService) VerifyRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, common.ErrorUnauthorized.Clone().WrapMessage("Invalid signing method")
		}
		return []byte(as.refreshSecret), nil
	})

	if err != nil {
		return "", common.ErrorUnauthorized.Clone().WrapError(err)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", common.ErrorUnauthorized.Clone().WrapMessage("Invalid token")
	}

	// Get email from claims
	email, ok := claims["email"].(string)
	if !ok {
		return "", common.ErrorUnauthorized.Clone().WrapMessage("Invalid token claims")
	}

	return email, nil
}

func (as *authService) RefreshAccessToken(ctx context.Context, refreshToken string) (*models.Auth, error) {
	// Verify refresh token
	email, err := as.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Verify user still exists
	_, err = as.es.FindByEmail(ctx, email)
	if err != nil {
		return nil, common.ErrorUnauthorized.Clone().WrapMessage("User not found")
	}

	// Generate new tokens
	auth, err := as.GenerateTokens(email)
	if err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapError(err)
	}

	return auth, nil
}
