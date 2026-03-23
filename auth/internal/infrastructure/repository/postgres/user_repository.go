package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/domain"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository constructs a Postgres-backed user repository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create inserts a new user into storage.
func (repository *UserRepository) Create(ctx context.Context, user domain.User) error {
	const query = `
		INSERT INTO users (uuid, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := repository.db.ExecContext(
		ctx,
		query,
		user.Uuid,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}

// Update persists mutable fields of an existing user.
func (repository *UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = `
		UPDATE users
		SET email = $2,
		    password = $3,
		    updated_at = $4
		WHERE uuid = $1
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		user.Uuid,
		user.Email,
		user.Password,
		user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("reading updated user rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}

// Find returns a user by UUID or nil when the user does not exist.
func (repository *UserRepository) Find(ctx context.Context, userUuid uuid.UUID) (*domain.User, error) {
	const query = `
		SELECT uuid, email, password, created_at, updated_at
		FROM users
		WHERE uuid = $1
	`

	var user domain.User

	if err := repository.db.QueryRowContext(ctx, query, userUuid).Scan(
		&user.Uuid,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("finding user by uuid: %w", err)
	}

	return &user, nil
}

// FindByEmail returns a user by email or nil when the user does not exist.
func (repository *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
		SELECT uuid, email, password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user domain.User

	if err := repository.db.QueryRowContext(ctx, query, email).Scan(
		&user.Uuid,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("finding user by email: %w", err)
	}

	return &user, nil
}
