package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application/login_user"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/validation"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/validation/rules"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Login handles user authentication requests.
func (handler *Handler) Login(request *http.Request) (int, any) {
	var payload loginRequest

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("decoding login request: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"email":    payload.Email,
		"password": payload.Password,
	}, map[string][]rules.Rule{
		"email":    {rules.Required(), rules.Max(255)},
		"password": {rules.Required(), rules.Max(255)},
	})

	if err := validator.Validate(); err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusBadRequest, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("validating login request: %w", err)
	}

	result, err := handler.loginUserHandler.Handle(request.Context(), login_user.Command{
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		if errors.Is(err, login_user.ErrInvalidCredentials) {
			validationError = validation.NewError()
			validationError.AddViolation("email", "Invalid credentials")
			validationError.AddViolation("password", "Invalid credentials")
			return http.StatusUnauthorized, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("logging in user: %w", err)
	}

	return http.StatusOK, LoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}
