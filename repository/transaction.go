package repository

import (
	"context"
	"database/sql"

	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/sirupsen/logrus"
)

var (
	QueryInsertNewRecord = `INSERT INTO "transactions" ("datetime", "amount") VALUES ($1, $2);`
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) RecordTransaction(ctx context.Context, transaction model.Transaction) error {
	logrus.Info("start inserting new transaction record")
	_, err := r.db.ExecContext(ctx, QueryInsertNewRecord, transaction.Datetime, transaction.Amount)
	if err != nil {
		logrus.Errorf("r.db.ExecContext return an error when trying inserting new transaction; err=%v", err)
		return err
	}

	logrus.Info("success inserting new transaction record")
	return nil
}
