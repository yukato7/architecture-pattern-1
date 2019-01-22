package mysql

import (
	"context"
	"database/sql"
	"github.com/yutify/gengo-api/domain/model"
	sq "gopkg.in/Masterminds/squirrel.v1"
	"time"
)

func NewChargeRepository(cm *Client, cs *Client) ChargeRepository {
	return &chargeRepository{
		DBMClient: cm,
		DBSClient: cs,
	}
}

type chargeRepository struct {
	DBMClient *Client
	DBSClient *Client
}

type ChargeRepository interface {
	CreateChargeLog(ctx context.Context, detail *model.ChargeDetail) error
}

func (r *chargeRepository) CreateChargeLog(ctx context.Context, detail *model.ChargeDetail) error {
	now := time.Now().Unix()
	query := sq.Insert(chargeLogTable).SetMap(sq.Eq{
		"user_id":    detail.UserID,
		"amount":     detail.Amount,
		"currency":   detail.Currency,
		"created_at": now,
	})
	return withTransaction(ctx, r.DBMClient.DB, func(tx *sql.Tx) error {
		_, err := query.RunWith(tx).ExecContext(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
