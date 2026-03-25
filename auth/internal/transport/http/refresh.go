package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/application/refresh_access_token"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
	"net/http"
)

type refreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

// Refresh reissues an access token using a refresh token.
func (handler *Handler) Refresh(request *http.Request) (int, any) {
	var payload refreshRequest

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("decoding refresh request: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"refreshToken": payload.RefreshToken,
	}, map[string][]rules.Rule{
		"refreshToken": {rules.Required()},
	})

	if err := validator.Validate(); err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusBadRequest, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("validating refresh request: %w", err)
	}

	result, err := handler.refreshAccessTokenHandler.Handle(request.Context(), refresh_access_token.Command{
		RefreshToken: payload.RefreshToken,
	})
	if err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		if errors.Is(err, refresh_access_token.ErrInvalidRefreshToken) {
			validationError = validation.NewError()
			validationError.AddViolation("refreshToken", "Invalid refresh token")
			return http.StatusUnauthorized, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("refreshing access token: %w", err)
	}

	return http.StatusOK, RefreshResponse{
		AccessToken: result.AccessToken,
	}
}
