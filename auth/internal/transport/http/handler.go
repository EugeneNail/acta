package http

import (
	"github.com/EugeneNail/acta/auth/internal/application/create_user"
	"github.com/EugeneNail/acta/auth/internal/application/login_user"
	"github.com/EugeneNail/acta/auth/internal/application/refresh_access_token"
)

type Handler struct {
	createUserHandler         *create_user.Handler
	loginUserHandler          *login_user.Handler
	refreshAccessTokenHandler *refresh_access_token.Handler
}

// NewHandler constructs the HTTP handler set.
func NewHandler(
	createUserHandler *create_user.Handler,
	loginUserHandler *login_user.Handler,
	refreshAccessTokenHandler *refresh_access_token.Handler,
) *Handler {
	return &Handler{
		createUserHandler:         createUserHandler,
		loginUserHandler:          loginUserHandler,
		refreshAccessTokenHandler: refreshAccessTokenHandler,
	}
}
