package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/EugeneNail/acta/journal/internal/domain"
	"github.com/google/uuid"
)

type HabitRepository struct {
	db *sql.DB
}

// NewHabitRepository constructs a Postgres-backed habit repository.
func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{
		db: db,
	}
}

// Create inserts a new habit into storage.
func (repository *HabitRepository) Create(ctx context.Context, habit domain.Habit) error {
	const query = `
		INSERT INTO habits (uuid, user_uuid, icon, name, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	if _, err := repository.db.ExecContext(
		ctx,
		query,
		habit.Uuid,
		habit.UserUuid,
		habit.Icon,
		habit.Name,
		habit.CreatedAt,
		habit.UpdatedAt,
		habit.DeletedAt,
	); err != nil {
		return fmt.Errorf("creating habit '%s' for user '%s': %w", habit.Uuid, habit.UserUuid, err)
	}

	return nil
}

// Update persists mutable fields of an existing habit.
func (repository *HabitRepository) Update(ctx context.Context, habit domain.Habit) error {
	const query = `
		UPDATE habits
		SET icon = $2,
		    name = $3,
		    updated_at = $4,
		    deleted_at = $5
		WHERE uuid = $1
	`

	if _, err := repository.db.ExecContext(
		ctx,
		query,
		habit.Uuid,
		habit.Icon,
		habit.Name,
		habit.UpdatedAt,
		habit.DeletedAt,
	); err != nil {
		return fmt.Errorf("updating habit '%s' for user '%s': %w", habit.Uuid, habit.UserUuid, err)
	}

	return nil
}

// Find returns a habit by UUID or nil when the habit does not exist.
func (repository *HabitRepository) Find(ctx context.Context, habitUuid uuid.UUID) (*domain.Habit, error) {
	const query = `
		SELECT uuid, user_uuid, icon, name, created_at, updated_at, deleted_at
		FROM habits
		WHERE uuid = $1
	`

	var habit domain.Habit
	var icon int

	if err := repository.db.QueryRowContext(ctx, query, habitUuid).Scan(
		&habit.Uuid,
		&habit.UserUuid,
		&icon,
		&habit.Name,
		&habit.CreatedAt,
		&habit.UpdatedAt,
		&habit.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("finding habit '%s': %w", habitUuid, err)
	}

	habit.Icon = domain.Icon(icon)

	return &habit, nil
}

// List returns all non-deleted habits for the user.
func (repository *HabitRepository) List(ctx context.Context, userUuid uuid.UUID) ([]domain.Habit, error) {
	const query = `
		SELECT uuid, user_uuid, icon, name, created_at, updated_at, deleted_at
		FROM habits
		WHERE user_uuid = $1
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := repository.db.QueryContext(ctx, query, userUuid)
	if err != nil {
		return nil, fmt.Errorf("listing habits for user '%s': %w", userUuid, err)
	}
	defer rows.Close()

	habits := []domain.Habit{}

	for rows.Next() {
		var habit domain.Habit
		var icon int

		if err := rows.Scan(
			&habit.Uuid,
			&habit.UserUuid,
			&icon,
			&habit.Name,
			&habit.CreatedAt,
			&habit.UpdatedAt,
			&habit.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("reading listed habits for user '%s': %w", userUuid, err)
		}

		habit.Icon = domain.Icon(icon)
		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating listed habits for user '%s': %w", userUuid, err)
	}

	return habits, nil
}
