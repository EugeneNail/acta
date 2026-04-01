package update_habit

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
	"github.com/google/uuid"
)

var ErrHabitNotFound = errors.New("habit not found")

type Command struct {
	Uuid     uuid.UUID
	Icon     domain.Icon
	Name     string
	UserUuid uuid.UUID
}

type Handler struct {
	repository domain.HabitRepository
}

// NewHandler constructs the update habit use-case handler.
func NewHandler(repository domain.HabitRepository) *Handler {
	return &Handler{repository: repository}
}

// Handle updates an existing habit.
func (handler *Handler) Handle(ctx context.Context, command Command) (domain.Habit, error) {
	validator := validation.NewValidator(map[string]any{
		"icon": int(command.Icon),
		"name": command.Name,
	}, map[string][]rules.Rule{
		"icon": {rules.Required(), rules.Min(1), rules.Max(1000)},
		"name": {rules.Required(), rules.Regex(rules.SentencePattern), rules.Min(3), rules.Max(50)},
	})

	if err := validator.Validate(); err != nil {
		return domain.Habit{}, fmt.Errorf("validating update habit command for habit '%s' and user '%s': %w", command.Uuid, command.UserUuid, err)
	}

	habit, err := handler.repository.Find(ctx, command.Uuid)
	if err != nil {
		return domain.Habit{}, fmt.Errorf("finding habit '%s' for user '%s': %w", command.Uuid, command.UserUuid, err)
	}

	if habit == nil || habit.UserUuid != command.UserUuid || habit.DeletedAt != nil {
		return domain.Habit{}, ErrHabitNotFound
	}

	habit.Icon = command.Icon
	habit.Name = command.Name
	habit.UpdatedAt = time.Now()

	if err := handler.repository.Update(ctx, *habit); err != nil {
		return domain.Habit{}, fmt.Errorf("updating habit '%s' for user '%s': %w", habit.Uuid, habit.UserUuid, err)
	}

	return *habit, nil
}
