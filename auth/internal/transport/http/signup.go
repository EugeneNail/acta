package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application/create_user"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/validation"
	"github.com/EugeneNail/acta/auth/internal/infrastructure/validation/rules"
	"github.com/google/uuid"
	"net/http"
)

type signupRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type SignupResponse struct {
	Uuid uuid.UUID `json:"uuid"`
}

// Signup handles user registration requests.
func (handler *Handler) Signup(request *http.Request) (int, any) {
	var payload signupRequest

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("decoding signup request: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"email":                payload.Email,
		"password":             payload.Password,
		"passwordConfirmation": payload.PasswordConfirmation,
	}, map[string][]rules.Rule{
		"email":                {rules.Required(), rules.Max(255)},
		"password":             {rules.Required(), rules.Max(255)},
		"passwordConfirmation": {rules.Required(), rules.Max(255)},
	})

	if err := validator.Validate(); err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusBadRequest, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("validating signup request: %w", err)
	}

	userUuid, err := handler.createUserHandler.Handle(request.Context(), create_user.Command{
		Email:                payload.Email,
		Password:             payload.Password,
		PasswordConfirmation: payload.PasswordConfirmation,
	})

	if err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		if errors.Is(err, create_user.ErrPasswordsDoNotMatch) {
			validationError = validation.NewError()
			validationError.AddViolation("password", "Passwords do not match")
			validationError.AddViolation("passwordConfirmation", "Passwords do not match")
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		if errors.Is(err, create_user.ErrUserAlreadyExists) {
			validationError = validation.NewError()
			validationError.AddViolation("email", "Email is already taken")
			return http.StatusConflict, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("registering user: %w", err)
	}

	return http.StatusCreated, SignupResponse{Uuid: userUuid}
}
