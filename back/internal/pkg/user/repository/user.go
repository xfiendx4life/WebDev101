package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/xfiendx4life/web101/meet/internal/models"
)

// CRUD Interface for user
type UserRepository interface {
	AddUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error)
	GetUserByName(ctx context.Context, name string) (models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}