package authentication

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const tempPublicKey = "TEMP_PRIVATE_KEY"
const tokenIssuer = "auth"
const audience = "acta"

type contextKey string

const userUUIDContextKey contextKey = "userUUID"

type claims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// WithUserUUID stores the authenticated user UUID in context.
func WithUserUUID(ctx context.Context, userUUID uuid.UUID) context.Context {
	return context.WithValue(ctx, userUUIDContextKey, userUUID)
}

// UserUUIDFromContext returns the authenticated user UUID from context.
func UserUUIDFromContext(ctx context.Context) uuid.UUID {
	userUUID, _ := ctx.Value(userUUIDContextKey).(uuid.UUID)

	return userUUID
}

// ParseAccessToken validates an access token and returns its subject user UUID.
func ParseAccessToken(token string) (uuid.UUID, error) {
	tokenClaims := claims{}
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&tokenClaims,
		func(parsedToken *jwt.Token) (any, error) {
			if parsedToken.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("reading signing method: unexpected method %q", parsedToken.Method.Alg())
			}

			return []byte(tempPublicKey), nil
		},
		jwt.WithIssuer(tokenIssuer),
		jwt.WithAudience(audience),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parsing access token: %w", err)
	}

	if !parsedToken.Valid {
		return uuid.Nil, errors.New("the token is invalid")
	}

	if tokenClaims.Type != "access" {
		return uuid.Nil, fmt.Errorf("reading token type: token type %q is not supported", tokenClaims.Type)
	}

	userUUID, err := uuid.Parse(tokenClaims.Subject)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parsing token subject: %w", err)
	}

	return userUUID, nil
}
