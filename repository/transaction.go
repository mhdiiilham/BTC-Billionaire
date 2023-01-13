package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/sirupsen/logrus"
)

var (
	QueryInsertNewRecord = `INSERT INTO "transactions" ("datetime", "amount") VALUES ($1, $2);`
	QueryUpdateBalance   = `
	INSERT INTO balance_histories (amount, datetime)
	VALUES ($1, $2)
	ON CONFLICT (datetime) DO update SET amount = (
		SELECT amount + $1 FROM balance_histories WHERE datetime = $2
	);`
	QueryGetBalanceHistory = `SELECT datetime, amount FROM balance_histories WHERE datetime >= $1 AND datetime <= $2;`
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

func (r *TransactionRepository) UpdateHourlyBalance(ctx context.Context, transaction model.Transaction) error {
	hourlyDate := transaction.Datetime.Format("2006-01-02 15:00:00")
	_, err := r.db.ExecContext(ctx, QueryUpdateBalance, transaction.Amount, hourlyDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) GetBalanceHistory(ctx context.Context, from, to string) ([]model.BalanceHistory, error) {
	result := []model.BalanceHistory{}

	rows, err := r.db.QueryContext(ctx, QueryGetBalanceHistory, from, to)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}

		return nil, err
	}

	for rows.Next() {
		var trx model.BalanceHistory
		rows.Scan(&trx.Datetime, &trx.Amount)
		result = append(result, trx)

	}

	return result, nil
}
