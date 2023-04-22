package postgres_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xfiendx4life/web101/meet/internal/models"
	"github.com/xfiendx4life/web101/meet/internal/pkg/storage"
	"github.com/xfiendx4life/web101/meet/internal/pkg/user/repository/postgres"
)

var store, _ = storage.New("postgresql://localhost:5433/sport?user=web101&password=123")

func TestAddUser(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	err := userRepo.AddUser(context.Background(), &user)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	res := store.Pool.QueryRow(context.Background(), "SELECT id, name, password, bio from users WHERE id=$1", user.ID)
	require.NoError(t, err)
	testUserResult := models.User{}
	err = res.Scan(&testUserResult.ID, &testUserResult.Name, &testUserResult.Password, &testUserResult.BIO)
	assert.NoError(t, err)
	assert.Equal(t, "test", testUserResult.Name)
	assert.Equal(t, "123", testUserResult.Password)
	assert.Equal(t, "Hello, this is the test user", testUserResult.BIO)
}

func TestAddUserError(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	store.Pool.Exec(context.Background(), `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
	err := userRepo.AddUser(context.Background(), &user)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
}

func TestAddUserErrorContext(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := userRepo.AddUser(ctx, &user)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
	assert.Equal(t, "context closed", err.Error())
}

func TestGetUser(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	row := store.Pool.QueryRow(context.Background(), `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
	var id uuid.UUID
	row.Scan(&id)
	res, err := userRepo.GetUserByID(context.Background(), id)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
	assert.Equal(t, "test", res.Name)
	assert.Equal(t, "123", res.Password)
	assert.Equal(t, "Hello, this is the test user", res.BIO)
}

func TestGetUserError(t *testing.T) {
	userRepo := postgres.New(store)
	id, _ := uuid.NewRandom()
	_, err := userRepo.GetUserByID(context.Background(), id)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
}

func TestGetUserErrorContext(t *testing.T) {
	userRepo := postgres.New(store)
	id, _ := uuid.NewRandom()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := userRepo.GetUserByID(ctx, id)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
	assert.ErrorIs(t, postgres.ErrClosedContext, err)
}

func TestGetUserByName(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	row := store.Pool.QueryRow(context.Background(), `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
	var id uuid.UUID
	row.Scan(&id)
	res, err := userRepo.GetUserByName(context.Background(), user.Name)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.NoError(t, err)
	assert.Equal(t, id, res.ID)
	assert.Equal(t, "123", res.Password)
	assert.Equal(t, "Hello, this is the test user", res.BIO)
}

func TestGetUserByNameError(t *testing.T) {
	userRepo := postgres.New(store)
	_, err := userRepo.GetUserByName(context.Background(), "test")
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
}

func TestGetUserByNameErrorContext(t *testing.T) {
	userRepo := postgres.New(store)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := userRepo.GetUserByName(ctx, "test")
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
	assert.ErrorIs(t, postgres.ErrClosedContext, err)
}

func TestUpdateUser(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	row := store.Pool.QueryRow(context.Background(), `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
	row.Scan(&user.ID)
	user.Name = "changedTest"
	user.BIO = "changedBIO"
	err := userRepo.UpdateUser(context.Background(), &user)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	res := store.Pool.QueryRow(context.Background(), "SELECT id, name, password, bio from users WHERE id=$1", user.ID)
	resUser := models.User{}
	res.Scan(&resUser.ID, &resUser.Name, &resUser.Password, &resUser.BIO)
	require.NoError(t, err)
	assert.Equal(t, user.ID, resUser.ID)
	assert.Equal(t, user.Name, resUser.Name)
	assert.Equal(t, user.BIO, resUser.BIO)
}

func TestUpdateError(t *testing.T) {
	userRepo := postgres.New(store)
	err := userRepo.UpdateUser(context.Background(), &models.User{ID: uuid.Nil})
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
}

func TestUpdateErrorContext(t *testing.T) {
	userRepo := postgres.New(store)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := userRepo.UpdateUser(ctx, &models.User{})
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	require.Error(t, err)
	assert.ErrorIs(t, postgres.ErrClosedContext, err)
}

func TestDeleteUser(t *testing.T) {
	userRepo := postgres.New(store)
	user := models.User{
		Name:     "test",
		Password: "123",
		BIO:      "Hello, this is the test user",
	}
	row := store.Pool.QueryRow(context.Background(), `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
	row.Scan(&user.ID)
	
	err := userRepo.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
	res := store.Pool.QueryRow(context.Background(), "SELECT COUNT(id) from users WHERE id=$1", user.ID)
	var count int
	res.Scan(&count)
	assert.Equal(t, 0, count)
}

func TestDeleteUserNoUser(t *testing.T) {
	userRepo := postgres.New(store)
	err := userRepo.DeleteUser(context.Background(), uuid.New())
	require.NoError(t, err)
	defer store.Pool.Exec(context.Background(), "DELETE FROM users")
}
