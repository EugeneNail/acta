package http

import (
	"github.com/EugeneNail/acta/journal/internal/application/create_habit"
)

type Handler struct {
	createHabitHandler *create_habit.Handler
}

// NewHandler constructs the HTTP handler set.
func NewHandler(
	createHabitHandler *create_habit.Handler,
) *Handler {
	return &Handler{
		createHabitHandler: createHabitHandler,
	}
}
