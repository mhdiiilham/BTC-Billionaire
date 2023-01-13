package service

import (
	"context"
	"errors"
	"time"

	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	transactionRepository TransactionRepository
}

var (
	ErrInvalidDateTimeFormat = errors.New("invalid datetime format")
)

func NewTransactionService(transactionRepository TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
	}
}

func (s *TransactionService) RecordNewTransaction(ctx context.Context, dateTime string, amount float64) error {
	logrus.Info("start record new transaction")

	trxTime, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		logrus.Errorf("failed parsing string time %s to time.Time due to error: %v", dateTime, err)
		return ErrInvalidDateTimeFormat
	}

	trx := model.Transaction{Datetime: trxTime, Amount: amount}
	if err := s.transactionRepository.RecordTransaction(ctx, trx); err != nil {
		logrus.Errorf("s.transactionRepository.RecordTransaction return an error %v", err)
		return err
	}

	logrus.Info("success record new transaction")
	return nil
}