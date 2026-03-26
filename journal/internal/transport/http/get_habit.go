package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/application/get_habit"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
	"github.com/EugeneNail/acta/journal/internal/transport/http/resource"
	"github.com/google/uuid"
)

// GetHabit handles habit lookup requests.
func (handler *Handler) GetHabit(request *http.Request) (int, any) {
	userUUID := authentication.UserUUIDFromContext(request.Context())

	habitUuid, err := uuid.Parse(request.PathValue("uuid"))
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("parsing habit uuid: %w", err)
	}

	habit, err := handler.getHabitHandler.Handle(request.Context(), get_habit.Query{
		Uuid:     habitUuid,
		UserUuid: userUUID,
	})
	if err != nil {
		if errors.Is(err, get_habit.ErrHabitNotFound) {
			return http.StatusNotFound, "habit not found"
		}

		return http.StatusInternalServerError, fmt.Errorf("getting habit '%s' for user '%s': %w", habitUuid, userUUID, err)
	}

	return http.StatusOK, resource.NewHabit(habit)
}
