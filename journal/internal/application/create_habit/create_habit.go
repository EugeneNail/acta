package create_habit

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

var ErrHabitLimitExceeded = errors.New("habit limit exceeded")

type Command struct {
	Icon     domain.Icon
	Name     string
	UserUuid uuid.UUID
}

const maxHabitsPerUser = 20

type Handler struct {
	repository domain.HabitRepository
}

// NewHandler constructs the create habit use-case handler.
func NewHandler(repository domain.HabitRepository) *Handler {
	return &Handler{repository: repository}
}

// Handle creates a new habit.
func (handler *Handler) Handle(ctx context.Context, command Command) (domain.Habit, error) {
	validator := validation.NewValidator(map[string]any{
		"icon": int(command.Icon),
		"name": command.Name,
	}, map[string][]rules.Rule{
		"icon": {rules.Required(), rules.Min(1), rules.Max(1000)},
		// TODO: Normalize habit names by trimming surrounding whitespace before validation and persistence.
		"name": {rules.Required(), rules.Regex(rules.SentencePattern), rules.Min(3), rules.Max(50)},
	})

	if err := validator.Validate(); err != nil {
		return domain.Habit{}, fmt.Errorf("validating create habit command for user '%s': %w", command.UserUuid, err)
	}

	habits, err := handler.repository.List(ctx, command.UserUuid)
	if err != nil {
		return domain.Habit{}, fmt.Errorf("listing habits for user '%s' before creating a new habit: %w", command.UserUuid, err)
	}

	// TODO: Make this check atomic by using database-level synchronization.
	if len(habits) >= maxHabitsPerUser {
		return domain.Habit{}, fmt.Errorf("validating create habit command for user '%s': %w", command.UserUuid, ErrHabitLimitExceeded)
	}

	now := time.Now()
	habit := domain.Habit{
		Uuid:      uuid.New(),
		Icon:      command.Icon,
		Name:      command.Name,
		CreatedAt: now,
		UpdatedAt: now,
		UserUuid:  command.UserUuid,
	}

	if err := handler.repository.Create(ctx, habit); err != nil {
		return domain.Habit{}, fmt.Errorf("creating habit '%s' for user '%s': %w", habit.Uuid, habit.UserUuid, err)
	}

	return habit, nil
}
