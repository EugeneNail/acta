package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/application/create_habit"
	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
	"github.com/EugeneNail/acta/journal/internal/transport/http/resource"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
)

type createHabitRequest struct {
	Icon int    `json:"icon"`
	Name string `json:"name"`
}

// CreateHabit handles habit creation requests.
func (handler *Handler) CreateHabit(request *http.Request) (int, any) {
	userUUID := authentication.UserUUIDFromContext(request.Context())

	var payload createHabitRequest
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("decoding create habit request: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"icon": payload.Icon,
		"name": payload.Name,
	}, map[string][]rules.Rule{
		"icon": {rules.Required(), rules.Min(1)},
		"name": {rules.Required(), rules.Max(255)},
	})

	if err := validator.Validate(); err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusBadRequest, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("validating create habit request: %w", err)
	}

	habit, err := handler.createHabitHandler.Handle(request.Context(), create_habit.Command{
		Icon:     domain.Icon(payload.Icon),
		Name:     payload.Name,
		UserUuid: userUUID,
	})
	if err != nil {
		if errors.Is(err, create_habit.ErrHabitLimitExceeded) {
			return http.StatusConflict, err
		}

		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("creating habit for user '%s': %w", userUUID, err)
	}

	return http.StatusCreated, resource.NewHabit(habit)
}
