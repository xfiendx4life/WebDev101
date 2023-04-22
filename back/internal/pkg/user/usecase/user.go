package usecase

import (
	"context"

	"github.com/xfiendx4life/web101/meet/internal/models"
)

type UserUsecase interface {
	Register(ctx context.Context, user *models.User) error
	Authenticate(ctx context.Context, login, password string) (bool, error)
}
