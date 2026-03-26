package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/application/delete_habit"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
	"github.com/google/uuid"
)

// DeleteHabit handles habit deletion requests.
func (handler *Handler) DeleteHabit(request *http.Request) (int, any) {
	userUUID := authentication.UserUUIDFromContext(request.Context())

	habitUUID, err := uuid.Parse(request.PathValue("uuid"))
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("parsing habit uuid: %w", err)
	}

	if err := handler.deleteHabitHandler.Handle(request.Context(), delete_habit.Command{
		Uuid:     habitUUID,
		UserUuid: userUUID,
	}); err != nil {
		if errors.Is(err, delete_habit.ErrHabitNotFound) {
			return http.StatusNotFound, "habit not found"
		}

		return http.StatusInternalServerError, fmt.Errorf("deleting habit '%s' for user '%s': %w", habitUUID, userUUID, err)
	}

	return http.StatusNoContent, nil
}
