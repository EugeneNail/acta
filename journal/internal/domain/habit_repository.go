package domain

import (
	"context"

	"github.com/google/uuid"
)

type HabitRepository interface {
	Create(ctx context.Context, habit Habit) error
	Update(ctx context.Context, habit Habit) error
	Find(ctx context.Context, uuid uuid.UUID) (*Habit, error)
	List(ctx context.Context, userUuid uuid.UUID) ([]Habit, error)
}
