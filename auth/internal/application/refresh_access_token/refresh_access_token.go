package refresh_access_token

import (
	"context"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application"
	"github.com/EugeneNail/acta/auth/internal/domain"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")

type Command struct {
	RefreshToken string
}

type Result struct {
	AccessToken string
}

type Handler struct {
	repository    domain.UserRepository
	tokenProvider application.TokenProvider
}

// NewHandler constructs the refresh access token use-case handler.
func NewHandler(repository domain.UserRepository, tokenProvider application.TokenProvider) *Handler {
	return &Handler{
		repository:    repository,
		tokenProvider: tokenProvider,
	}
}

// Handle validates a refresh token and returns a new access token.
func (handler *Handler) Handle(ctx context.Context, command Command) (Result, error) {
	userUUID, err := handler.tokenProvider.ParseRefreshToken(command.RefreshToken)
	if err != nil {
		return Result{}, ErrInvalidRefreshToken
	}

	user, err := handler.repository.Find(ctx, userUUID)
	if err != nil {
		return Result{}, fmt.Errorf("finding user by uuid: %w", err)
	}

	if user == nil {
		return Result{}, ErrInvalidRefreshToken
	}

	accessToken, err := handler.tokenProvider.GenerateAccessToken(*user)
	if err != nil {
		return Result{}, fmt.Errorf("generating access token: %w", err)
	}

	return Result{
		AccessToken: accessToken,
	}, nil
}
