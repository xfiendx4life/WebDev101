package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/xfiendx4life/web101/meet/internal/models"
	"github.com/xfiendx4life/web101/meet/internal/pkg/storage"
	"github.com/xfiendx4life/web101/meet/internal/pkg/user/repository"
)

var ErrClosedContext = errors.New("context closed")

type userPg struct {
	store storage.Gres
}

func New(store storage.Gres) repository.UserRepository {
	return &userPg{
		store: store,
	}
}

func (upg *userPg) AddUser(ctx context.Context, user *models.User) error {
	select {
	case <-ctx.Done():
		return ErrClosedContext
	default:
		pool := upg.store.Pool
		row := pool.QueryRow(ctx, `INSERT INTO users (name, password, bio)
		VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Password, user.BIO)
		err := row.Scan(&user.ID)
		if err != nil {
			return fmt.Errorf("can't scan id: %w", err)
		}
		return nil
	}

}
func (upg *userPg) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (upg *userPg) GetUserByID(ctx context.Context, id uuid.UUID) (result models.User, err error) {
	select {
	case <-ctx.Done():
		err = ErrClosedContext
		return
	default:
		pool := upg.store.Pool
		row := pool.QueryRow(context.Background(), "SELECT id, name, password, bio from users WHERE id=$1", id)
		err = row.Scan(&result.ID, &result.Name, &result.Password, &result.BIO)
		if err != nil {
			return
		}
		return
	}
}
func (upg *userPg) GetUserByName(ctx context.Context, name string) (result models.User, err error) {
	select {
	case <-ctx.Done():
		err = ErrClosedContext
		return
	default:
		pool := upg.store.Pool
		row := pool.QueryRow(context.Background(), "SELECT id, name, password, bio from users WHERE name=$1", name)
		err = row.Scan(&result.ID, &result.Name, &result.Password, &result.BIO)
		if err != nil {
			return
		}
		return
	}
}
func (upg *userPg) UpdateUser(ctx context.Context, user *models.User) (err error) {
	select {
	case <-ctx.Done():
		err = ErrClosedContext
		return
	default:
		pool := upg.store.Pool
		var tag pgconn.CommandTag
		tag, err = pool.Exec(ctx, `UPDATE users SET name=$1, password=$2, bio=$3 WHERE id=$4`,
			user.Name, user.Password, user.BIO, user.ID)
		log.Println(tag)
		if err != nil || tag.String() == "UPDATE 0" {
			err = fmt.Errorf("can't update user %w", err)
		}
		return
	}
}
