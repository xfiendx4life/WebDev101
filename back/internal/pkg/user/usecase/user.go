package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xfiendx4life/web101/meet/internal/models"
)

var ErrorAuthentication = errors.New("failed authentication")

type UserUsecase interface {
	Register(ctx context.Context, user *models.User) error
	Authenticate(ctx context.Context, login, password string) (uuid.UUID, error)
}
