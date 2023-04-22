package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/xfiendx4life/web101/meet/internal/models"
	"github.com/xfiendx4life/web101/meet/internal/pkg/user/usecase"
)

type mockrepo struct {
	err error
}

func (mc *mockrepo) AddUser(ctx context.Context, user *models.User) error {
	return mc.err
}
func (mc *mockrepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return mc.err
}
func (mc *mockrepo) GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	return models.User{}, mc.err
}

func (mc *mockrepo) GetUserByName(ctx context.Context, name string) (models.User, error) {
	return testUser, mc.err
}
func (mc *mockrepo) UpdateUser(ctx context.Context, user *models.User) error {
	return mc.err
}

var testUser = models.User{
	Name:     "test",
	Password: usecase.HashPassword("test"),
}

func TestRegister(t *testing.T) {
	repo := mockrepo{}
	uc := usecase.New(&repo)
	err := uc.Register(context.Background(), &testUser)
	require.NoError(t, err)
}

func TestRegisterErrorDB(t *testing.T) {
	repo := mockrepo{err: errors.New("test error")}
	uc := usecase.New(&repo)
	err := uc.Register(context.Background(), &testUser)
	require.Error(t, err)
}

func TestRegisterError(t *testing.T) {
	repo := mockrepo{}
	uc := usecase.New(&repo)
	err := uc.Register(context.Background(), &models.User{})
	require.Error(t, err)
}

func TestAuth(t *testing.T) {
	repo := mockrepo{}
	uc := usecase.New(&repo)
	res, err := uc.Authenticate(context.Background(), "test", "test")
	require.NoError(t, err)
	require.True(t, res)
}

func TestAuthErrorDB(t *testing.T) {
	repo := mockrepo{err: errors.New("test error")}
	uc := usecase.New(&repo)
	res, err := uc.Authenticate(context.Background(), "test", "test")
	require.Error(t, err)
	require.False(t, res)
}

func TestAuthWrongPassw(t *testing.T) {
	repo := mockrepo{}
	uc := usecase.New(&repo)
	res, err := uc.Authenticate(context.Background(), "test", "ne test")
	require.NoError(t, err)
	require.False(t, res)
}

