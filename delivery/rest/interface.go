package rest

import (
	"context"

	"github.com/mhdiiilham/BTC-Billionaire/model"
)

type Transactioner interface {
	RecordNewTransaction(ctx context.Context, dateTime string, amount float64) error
	GetBalanceHistory(ctx context.Context, fromDate, toDate string) (balanceHistories []model.BalanceHistory, err error)
}
