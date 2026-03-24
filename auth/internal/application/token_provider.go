package application

import (
	"github.com/EugeneNail/acta/auth/internal/domain"
	"github.com/google/uuid"
)

type TokenProvider interface {
	GenerateAccessToken(user domain.User) (string, error)
	GenerateRefreshToken(user domain.User) (string, error)
	ParseRefreshToken(token string) (uuid.UUID, error)
}
