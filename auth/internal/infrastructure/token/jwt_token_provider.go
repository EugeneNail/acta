package token

import (
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const tempPrivateKey = "TEMP_PRIVATE_KEY"
const tokenIssuer = "auth"
const audience = "acta"
const accessTokenLifetime = 15 * time.Minute
const refreshTokenLifetime = 7 * 24 * time.Hour

type Claims struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

type Provider struct{}

// NewProvider constructs a JWT token provider.
func NewProvider() *Provider {
	return &Provider{}
}

// GenerateAccessToken creates a signed access token for the user.
func (provider *Provider) GenerateAccessToken(user domain.User) (string, error) {
	return provider.generateToken(user, "access", accessTokenLifetime)
}

// GenerateRefreshToken creates a signed refresh token for the user.
func (provider *Provider) GenerateRefreshToken(user domain.User) (string, error) {
	return provider.generateToken(user, "refresh", refreshTokenLifetime)
}

// ParseRefreshToken validates a refresh token and returns its subject user UUID.
func (provider *Provider) ParseRefreshToken(token string) (uuid.UUID, error) {
	claims := Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(parsedToken *jwt.Token) (any, error) {
		if parsedToken.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("reading signing method: unexpected method %q", parsedToken.Method.Alg())
		}

		return []byte(tempPrivateKey), nil
	}, jwt.WithIssuer(tokenIssuer), jwt.WithAudience(audience), jwt.WithExpirationRequired(), jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return uuid.Nil, fmt.Errorf("parsing refresh token: %w", err)
	}

	if !parsedToken.Valid {
		return uuid.Nil, errors.New("the token is invalid")
	}

	if claims.Type != "refresh" {
		return uuid.Nil, fmt.Errorf("reading token type: token type %q is not supported", claims.Type)
	}

	userUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parsing token subject: %w", err)
	}

	return userUUID, nil
}

// generateToken creates a signed JWT for the user with the provided type and lifetime.
func (provider *Provider) generateToken(user domain.User, tokenType string, lifetime time.Duration) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: user.Email,
		Type:  tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Uuid.String(),
			Issuer:    tokenIssuer,
			Audience:  []string{audience},
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(lifetime)),
		},
	})

	signedToken, err := token.SignedString([]byte(tempPrivateKey))
	if err != nil {
		return "", fmt.Errorf("signing %s token: %w", tokenType, err)
	}

	return signedToken, nil
}
