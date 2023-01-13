package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/mhdiiilham/BTC-Billionaire/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type transactionRepositoryTestSuite struct {
	suite.Suite
	db     *sql.DB
	dbMock sqlmock.Sqlmock
}

func TestTransactionRepository(t *testing.T) {
	suite.Run(t, new(transactionRepositoryTestSuite))
}

func (suite *transactionRepositoryTestSuite) TestRecordTransaction() {
	trxTime, err := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
	if err != nil {
		suite.FailNowf("failed to convert string time to time.Time", "error: %v", err)
	}
	amount := float64(10)

	testCases := []struct {
		name        string
		transaction model.Transaction
		doMocks     func()
		expected    error
	}{
		{
			name:        "success insert to db",
			transaction: model.Transaction{Datetime: trxTime, Amount: amount},
			doMocks: func() {
				suite.dbMock.
					ExpectExec(regexp.QuoteMeta(repository.QueryInsertNewRecord)).
					WithArgs(trxTime, amount).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name:        "failed insert to db",
			transaction: model.Transaction{Datetime: trxTime, Amount: amount},
			doMocks: func() {
				suite.dbMock.
					ExpectExec(regexp.QuoteMeta(repository.QueryInsertNewRecord)).
					WithArgs(trxTime, amount).
					WillReturnError(sql.ErrConnDone)
			},
			expected: sql.ErrConnDone,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tt.doMocks()

			transactionRepository := repository.NewTransactionRepository(suite.db)
			actual := transactionRepository.RecordTransaction(ctx, tt.transaction)
			assert.Equal(t, tt.expected, actual, "transactionRepository.RecordTransaction should return %v instead of %v", tt.expected, actual)
		})
	}
}

func (suite *transactionRepositoryTestSuite) TestUpdateHourlyBalance() {
	trxTime, err := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
	if err != nil {
		suite.FailNowf("failed to convert string time to time.Time", "error: %v", err)
	}
	amount := float64(10)

	testCases := []struct {
		name        string
		transaction model.Transaction
		doMocks     func()
		expected    error
	}{
		{
			name:        "success",
			transaction: model.Transaction{Datetime: trxTime, Amount: amount},
			doMocks: func() {
				hourlyDate := trxTime.Format("2006-01-02 15:00:00")
				suite.dbMock.
					ExpectExec(regexp.QuoteMeta(repository.QueryUpdateBalance)).
					WithArgs(amount, hourlyDate).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name:        "failed",
			transaction: model.Transaction{Datetime: trxTime, Amount: amount},
			doMocks: func() {
				hourlyDate := trxTime.Format("2006-01-02 15:00:00")
				suite.dbMock.
					ExpectExec(regexp.QuoteMeta(repository.QueryUpdateBalance)).
					WithArgs(amount, hourlyDate).
					WillReturnError(sql.ErrConnDone)
			},
			expected: sql.ErrConnDone,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.doMocks()
			ctx := context.Background()

			transactionRepository := repository.NewTransactionRepository(suite.db)
			actual := transactionRepository.UpdateHourlyBalance(ctx, tt.transaction)
			assert.Equal(t, tt.expected, actual, "transactionRepository.UpdateHourlyBalance should return %v instead of %v", tt.expected, actual)
		})
	}
}

func (suite *transactionRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.FailNowf("failed setup test", "failed get new instance of sqlmock: %v", err)
	}

	suite.db = db
	suite.dbMock = mock
}

func (suite *transactionRepositoryTestSuite) TearDownTest() {
	defer suite.db.Close()
}
