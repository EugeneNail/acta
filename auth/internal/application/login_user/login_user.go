package login_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application"
	"github.com/EugeneNail/acta/auth/internal/domain"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Command struct {
	Email    string
	Password string
}

type Result struct {
	AccessToken  string
	RefreshToken string
}

type Handler struct {
	repository    domain.UserRepository
	tokenProvider application.TokenProvider
}

// NewHandler constructs the login user use-case handler.
func NewHandler(repository domain.UserRepository, tokenProvider application.TokenProvider) *Handler {
	return &Handler{
		repository:    repository,
		tokenProvider: tokenProvider,
	}
}

// Handle authenticates a user and returns a new access/refresh token pair.
func (handler *Handler) Handle(ctx context.Context, command Command) (Result, error) {
	validator := validation.NewValidator(map[string]any{
		"email":    command.Email,
		"password": command.Password,
	}, map[string][]rules.Rule{
		"email":    {rules.Required(), rules.Min(3), rules.Max(100), rules.Regex(rules.EmailPattern)},
		"password": {rules.Required(), rules.Min(8), rules.Max(100)},
	})

	if err := validator.Validate(); err != nil {
		return Result{}, fmt.Errorf("validating login command: %w", err)
	}

	user, err := handler.repository.FindByEmail(ctx, command.Email)
	if err != nil {
		return Result{}, fmt.Errorf("finding user by email: %w", err)
	}

	if user == nil {
		return Result{}, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.Password)); err != nil {
		return Result{}, ErrInvalidCredentials
	}

	accessToken, err := handler.tokenProvider.GenerateAccessToken(*user)
	if err != nil {
		return Result{}, fmt.Errorf("generating access token: %w", err)
	}

	refreshToken, err := handler.tokenProvider.GenerateRefreshToken(*user)
	if err != nil {
		return Result{}, fmt.Errorf("generating refresh token: %w", err)
	}

	return Result{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
