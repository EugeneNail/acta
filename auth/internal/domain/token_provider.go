package domain

import "github.com/google/uuid"

type TokenProvider interface {
	GenerateAccessToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	ParseRefreshToken(token string) (uuid.UUID, error)
}
