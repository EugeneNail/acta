package http

import (
	"fmt"
	"net/http"

	"github.com/EugeneNail/acta/journal/internal/application/list_habits"
	"github.com/EugeneNail/acta/journal/internal/infrastructure/authentication"
	"github.com/EugeneNail/acta/journal/internal/transport/http/resource"
)

// ListHabits handles habit list requests.
func (handler *Handler) ListHabits(request *http.Request) (int, any) {
	userUUID := authentication.UserUUIDFromContext(request.Context())

	habits, err := handler.listHabitsHandler.Handle(request.Context(), list_habits.Query{
		UserUuid: userUUID,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("listing habits for user '%s': %w", userUUID, err)
	}

	resources := make([]resource.Habit, 0, len(habits))
	for _, habit := range habits {
		resources = append(resources, resource.NewHabit(habit))
	}

	return http.StatusOK, resources
}
