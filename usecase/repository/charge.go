package repository

import (
	"context"
	"github.com/yutify/gengo-api/domain/model"
)

type ChargeRepository interface {
	CreateChargeLog(ctx context.Context, detail *model.ChargeDetail) error
}
