package domain

import (
	"github.com/google/uuid"
	"time"
)

type Habit struct {
	Uuid      uuid.UUID
	Icon      Icon
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserUuid uuid.UUID
}
