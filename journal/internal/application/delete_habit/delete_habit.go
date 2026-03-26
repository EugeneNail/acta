package delete_habit

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/google/uuid"
)

var ErrHabitNotFound = errors.New("habit not found")

type Command struct {
	Uuid     uuid.UUID
	UserUuid uuid.UUID
}

type Handler struct {
	repository domain.HabitRepository
}

// NewHandler constructs the delete habit use-case handler.
func NewHandler(repository domain.HabitRepository) *Handler {
	return &Handler{repository: repository}
}

// Handle soft-deletes a habit.
func (handler *Handler) Handle(ctx context.Context, command Command) error {
	habit, err := handler.repository.Find(ctx, command.Uuid)
	if err != nil {
		return fmt.Errorf("finding habit '%s' for user '%s': %w", command.Uuid, command.UserUuid, err)
	}

	if habit == nil || habit.UserUuid != command.UserUuid || habit.DeletedAt != nil {
		return ErrHabitNotFound
	}

	now := time.Now()
	habit.DeletedAt = &now
	habit.UpdatedAt = now

	if err := handler.repository.Update(ctx, *habit); err != nil {
		return fmt.Errorf("deleting habit '%s' for user '%s': %w", habit.Uuid, habit.UserUuid, err)
	}

	return nil
}
