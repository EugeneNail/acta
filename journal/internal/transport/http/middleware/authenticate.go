package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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

// Authenticate validates an access token and stores the user UUID in the request context.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authorizationHeader := request.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(writer, "authorization header is required", http.StatusUnauthorized)
			return
		}

		token, found := strings.CutPrefix(authorizationHeader, "Bearer ")
		if !found || token == "" {
			http.Error(writer, "authorization header must use Bearer token", http.StatusUnauthorized)
			return
		}

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
			http.Error(writer, fmt.Sprintf("parsing access token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		if !parsedToken.Valid {
			http.Error(writer, "the token is invalid", http.StatusUnauthorized)
			return
		}

		if tokenClaims.Type != "access" {
			http.Error(writer, fmt.Sprintf("reading token type: token type %q is not supported", tokenClaims.Type), http.StatusUnauthorized)
			return
		}

		userUUID, err := uuid.Parse(tokenClaims.Subject)
		if err != nil {
			http.Error(writer, fmt.Sprintf("parsing token subject: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		next(writer, request.WithContext(context.WithValue(request.Context(), userUUIDContextKey, userUUID)))
	}
}

// UserUUID returns the authenticated user UUID from context.
func UserUUID(ctx context.Context) uuid.UUID {
	userUUID, _ := ctx.Value(userUUIDContextKey).(uuid.UUID)

	return userUUID
}
