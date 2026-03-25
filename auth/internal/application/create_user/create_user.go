package create_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/EugeneNail/acta/auth/internal/domain"
	"github.com/EugeneNail/acta/lib-common/pkg/validation"
	"github.com/EugeneNail/acta/lib-common/pkg/validation/rules"
	"github.com/google/uuid"
	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var ErrPasswordsDoNotMatch = errors.New("passwords do not match")
var ErrUserAlreadyExists = errors.New("user already exists")

type Command struct {
	Email                string
	Password             string
	PasswordConfirmation string
}

type Handler struct {
	repository domain.UserRepository
}

func NewHandler(repository domain.UserRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (handler *Handler) Handle(ctx context.Context, command Command) (uuid.UUID, error) {
	validator := validation.NewValidator(map[string]any{
		"email":                command.Email,
		"password":             command.Password,
		"passwordConfirmation": command.PasswordConfirmation,
	}, map[string][]rules.Rule{
		"email":                {rules.Required(), rules.Min(3), rules.Max(100), rules.Regex(rules.EmailPattern)},
		"password":             {rules.Required(), rules.Min(8), rules.Max(100), rules.Password()},
		"passwordConfirmation": {rules.Required(), rules.Same("password")},
	})

	if err := validator.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("validating registration command: %w", err)
	}

	if command.Password != command.PasswordConfirmation {
		return uuid.Nil, ErrPasswordsDoNotMatch
	}

	existingUser, err := handler.repository.FindByEmail(ctx, command.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("finding user by email: %w", err)
	}

	if existingUser != nil {
		return uuid.Nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(command.Password), 12)
	if err != nil {
		return uuid.Nil, fmt.Errorf("hashing password: %w", err)
	}

	now := time.Now()
	user := domain.User{
		Uuid:      uuid.UUID(uuidv7.New()),
		Email:     command.Email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := handler.repository.Create(ctx, user); err != nil {
		return uuid.Nil, fmt.Errorf("creating user: %w", err)
	}

	return user.Uuid, nil
}
