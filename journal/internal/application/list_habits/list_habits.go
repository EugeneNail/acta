package list_habits

import (
	"context"
	"fmt"

	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/google/uuid"
)

type Query struct {
	UserUuid uuid.UUID
}

type Handler struct {
	repository domain.HabitRepository
}

// NewHandler constructs the list habits use-case handler.
func NewHandler(repository domain.HabitRepository) *Handler {
	return &Handler{repository: repository}
}

// Handle returns all habits for the user.
func (handler *Handler) Handle(ctx context.Context, query Query) ([]domain.Habit, error) {
	habits, err := handler.repository.List(ctx, query.UserUuid)
	if err != nil {
		return nil, fmt.Errorf("listing habits for user '%s': %w", query.UserUuid, err)
	}

	return habits, nil
}
