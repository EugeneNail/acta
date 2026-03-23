package domain

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user *User) error
	Find(ctx context.Context, uuid uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
