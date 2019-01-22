package mysql

import (
	"context"
	"database/sql"
	"github.com/yutify/gengo-api/domain/model"
	sq "gopkg.in/Masterminds/squirrel.v1"
	"time"
)

type userRepository struct {
	DBMClient *Client
	DBSClient *Client
}

func NewUserRepository(cm *Client, cs *Client) UserRepository {
	return &userRepository{
		DBMClient: cm,
		DBSClient: cs,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID string) (*model.User, error)
	UserExists(ctx context.Context, userID string) (bool, error)
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	now := time.Now().Unix()
	query := sq.Insert(userTable).SetMap(sq.Eq{ //Todo add Transaction
		"id":         user.ID,
		"name":       user.Name,
		"icon_url":   user.IconURL,
		"created_at": now,
		"updated_at": now,
	})
	return withTransaction(ctx, r.DBMClient.DB, func(tx *sql.Tx) error {
		_, err := query.RunWith(tx).ExecContext(ctx)
		if err != nil { //Todo add error handling
			return err
		}
		return nil
	})
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	query := sq.Select("name", "icon_url"). //Todo add Transaction
						From(userTable).
						Where(sq.Eq{
			"id": userID,
		})
	err := query.RunWith(r.DBSClient.DB).QueryRowContext(ctx).Scan(&user.Name, &user.IconURL)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) UserExists(ctx context.Context, userID string) (bool, error) {
	query := sq.Select("id").
		From(userTable).
		Where(sq.Eq{
			"id": userID,
		})
	rows, err := query.RunWith(r.DBSClient.DB).QueryContext(ctx)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}
