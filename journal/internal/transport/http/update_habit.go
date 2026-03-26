package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/application/update_habit"
	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
	"github.com/EugeneNail/acta/journal/internal/transport/http/resource"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
	"github.com/google/uuid"
)

type updateHabitRequest struct {
	Icon int    `json:"icon"`
	Name string `json:"name"`
}

// UpdateHabit handles habit update requests.
func (handler *Handler) UpdateHabit(request *http.Request) (int, any) {
	userUUID := authentication.UserUUIDFromContext(request.Context())

	var payload updateHabitRequest
	habitUuid, err := uuid.Parse(request.PathValue("uuid"))
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("parsing habit uuid: %w", err)
	}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("decoding update habit request: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"icon": payload.Icon,
		"name": payload.Name,
	}, map[string][]rules.Rule{
		"icon": {rules.Min(1)},
		"name": {rules.Required(), rules.Max(255)},
	})

	if err := validator.Validate(); err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusBadRequest, validationError.Violations()
		}

		return http.StatusInternalServerError, fmt.Errorf("validating update habit request: %w", err)
	}

	habit, err := handler.updateHabitHandler.Handle(request.Context(), update_habit.Command{
		Uuid:     habitUuid,
		Icon:     domain.Icon(payload.Icon),
		Name:     payload.Name,
		UserUuid: userUUID,
	})
	if err != nil {
		var validationError validation.Error
		if errors.As(err, &validationError) {
			return http.StatusUnprocessableEntity, validationError.Violations()
		}

		if errors.Is(err, update_habit.ErrHabitNotFound) {
			return http.StatusNotFound, "habit not found"
		}

		return http.StatusInternalServerError, fmt.Errorf("updating habit '%s' for user '%s': %w", habitUuid, userUUID, err)
	}

	return http.StatusOK, resource.NewHabit(habit)
}
