package get_habit

import (
	"context"
	"errors"
	"fmt"

	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/google/uuid"
)

var ErrHabitNotFound = errors.New("habit not found")

type Query struct {
	Uuid     uuid.UUID
	UserUuid uuid.UUID
}

type Handler struct {
	repository domain.HabitRepository
}

// NewHandler constructs the get habit use-case handler.
func NewHandler(repository domain.HabitRepository) *Handler {
	return &Handler{repository: repository}
}

// Handle returns a habit by UUID for the user.
func (handler *Handler) Handle(ctx context.Context, query Query) (domain.Habit, error) {
	habit, err := handler.repository.Find(ctx, query.Uuid)
	if err != nil {
		return domain.Habit{}, fmt.Errorf("finding habit '%s' for user '%s': %w", query.Uuid, query.UserUuid, err)
	}

	if habit == nil || habit.UserUuid != query.UserUuid || habit.DeletedAt != nil {
		return domain.Habit{}, ErrHabitNotFound
	}

	return *habit, nil
}
