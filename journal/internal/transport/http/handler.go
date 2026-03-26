package http

import (
	"github.com/EugeneNail/acta/journal/internal/application/create_habit"
	"github.com/EugeneNail/acta/journal/internal/application/delete_habit"
	"github.com/EugeneNail/acta/journal/internal/application/update_habit"
)

type Handler struct {
	createHabitHandler *create_habit.Handler
	deleteHabitHandler *delete_habit.Handler
	updateHabitHandler *update_habit.Handler
}

// NewHandler constructs the HTTP handler set.
func NewHandler(
	createHabitHandler *create_habit.Handler,
	deleteHabitHandler *delete_habit.Handler,
	updateHabitHandler *update_habit.Handler,
) *Handler {
	return &Handler{
		createHabitHandler: createHabitHandler,
		deleteHabitHandler: deleteHabitHandler,
		updateHabitHandler: updateHabitHandler,
	}
}
