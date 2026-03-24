package http

import (
	"github.com/EugeneNail/acta/auth/internal/application/create_user"
	"github.com/EugeneNail/acta/auth/internal/application/login_user"
)

type Handler struct {
	createUserHandler *create_user.Handler
	loginUserHandler  *login_user.Handler
}

// NewHandler constructs the HTTP handler set.
func NewHandler(
	createUserHandler *create_user.Handler,
	loginUserHandler *login_user.Handler,
) *Handler {
	return &Handler{
		createUserHandler: createUserHandler,
		loginUserHandler:  loginUserHandler,
	}
}
