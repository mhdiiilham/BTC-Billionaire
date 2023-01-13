package service

import (
	"context"

	"github.com/mhdiiilham/BTC-Billionaire/model"
)

type TransactionRepository interface {
	RecordTransaction(ctx context.Context, transaction model.Transaction) error
	UpdateHourlyBalance(ctx context.Context, transaction model.Transaction) error
	GetBalanceHistory(ctx context.Context, from, to string) ([]model.BalanceHistory, error)
}
