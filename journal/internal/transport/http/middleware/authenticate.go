package middleware

import (
	"net/http"
	"strings"

	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
)

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

		userUUID, err := authentication.ParseAccessToken(token)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		next(writer, request.WithContext(authentication.WithUserUUID(request.Context(), userUUID)))
	}
}
