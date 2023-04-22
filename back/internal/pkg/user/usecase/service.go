package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/xfiendx4life/web101/meet/internal/models"
	"github.com/xfiendx4life/web101/meet/internal/pkg/user/repository"
)

type usecase struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UserUsecase {
	return &usecase{
		repo: repo,
	}
}

var ErrorContextDoneUsecase = errors.New("end with context in usecase")
var _ UserUsecase = (*usecase)(nil)

func HashPassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return string(h.Sum(nil))
}

func (us *usecase) Register(ctx context.Context, user *models.User) error {
	select {
	case <-ctx.Done():
		return ErrorContextDoneUsecase
	default:
		if user.Name == "" || user.Password == "" {
			return errors.New("empty name or password")
		}
		user.Password = HashPassword(user.Password)
		err := us.repo.AddUser(ctx, user)
		if err != nil {
			return fmt.Errorf("can't register user: %w", err)
		}
		return nil
	}
}

func (us *usecase) Authenticate(ctx context.Context, login, password string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ErrorContextDoneUsecase
	default:
		if login == "" {
			return false, nil
		}
		user, err := us.repo.GetUserByName(ctx, login)
		if err != nil {
			return false, fmt.Errorf("can't authenticate %w", err)
		}
		return user.Password == HashPassword(password), nil

	}

}
