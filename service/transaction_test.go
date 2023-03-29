package service_test

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/mhdiiilham/BTC-Billionaire/service"
	"github.com/mhdiiilham/BTC-Billionaire/service/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type transactionTestsSuite struct {
	suite.Suite
	ctrl                      *gomock.Controller
	mockTransactionRepository *mock.MockTransactionRepository
}

func TestTransactionService(t *testing.T) {
	suite.Run(t, new(transactionTestsSuite))
}

func (suite *transactionTestsSuite) TestRecordNewTransaction() {
	testCases := []struct {
		name     string
		dateTime string
		amount   float64
		doMocks  func(wg *sync.WaitGroup)
		expected error
	}{
		{
			name:     "amount is less or equal than zero",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   0,
			doMocks: func(wg *sync.WaitGroup) {
				wg.Add(1)
				wg.Done()
			},
			expected: service.ErrInvalidAmount,
		},
		{
			name:     "failed to covert dateTime string to time.Time",
			dateTime: "january the third",
			amount:   10,
			doMocks: func(wg *sync.WaitGroup) {
				wg.Add(1)
				wg.Done()
			},
			expected: service.ErrInvalidDateTimeFormat,
		},
		{
			name:     "success record new transaction to database",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   10,
			doMocks: func(wg *sync.WaitGroup) {

				trxTime, _ := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
				trx := model.Transaction{
					Datetime: trxTime,
					Amount:   10,
				}

				suite.
					mockTransactionRepository.
					EXPECT().
					RecordTransaction(gomock.Any(), trx).
					Return(nil).
					Times(1)

				wg.Add(1)
				suite.
					mockTransactionRepository.
					EXPECT().
					UpdateHourlyBalance(gomock.Any(), trx).
					Return(nil).
					Do(func(arg0, arg1 interface{}) {
						defer wg.Done()
					}).
					Times(1)
			},
			expected: nil,
		},
		{
			name:     "failed record new transaction to database",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   10,
			doMocks: func(wg *sync.WaitGroup) {
				wg.Add(1)
				wg.Done()

				trxTime, _ := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
				suite.
					mockTransactionRepository.
					EXPECT().
					RecordTransaction(gomock.Any(), model.Transaction{
						Datetime: trxTime,
						Amount:   10,
					}).Return(sql.ErrConnDone).
					Times(1)
			},
			expected: sql.ErrConnDone,
		},
		{
			name:     "failed to update balance",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   10,
			doMocks: func(wg *sync.WaitGroup) {

				trxTime, _ := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
				trx := model.Transaction{
					Datetime: trxTime,
					Amount:   10,
				}

				suite.
					mockTransactionRepository.
					EXPECT().
					RecordTransaction(gomock.Any(), trx).
					Return(nil).
					Times(1)

				wg.Add(1)
				suite.
					mockTransactionRepository.
					EXPECT().
					UpdateHourlyBalance(gomock.Any(), trx).
					Return(sql.ErrConnDone).
					Do(func(arg0, arg1 interface{}) {
						defer wg.Done()
					}).
					Times(1)
			},
			expected: nil,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup

			ctx := context.Background()
			tt.doMocks(&wg)

			svc := service.NewTransactionService(suite.mockTransactionRepository)
			actual := svc.RecordNewTransaction(ctx, tt.dateTime, tt.amount)
			wg.Wait()

			assert.Equal(t, tt.expected, actual, "expected function NewTransactionRecord returning %v", tt.expected)
		})
	}
}

func (suite *transactionTestsSuite) TestGetBalanceHistory() {
	testCases := []struct {
		name        string
		fromDate    string
		toDate      string
		doMocks     func()
		expected    []model.BalanceHistory
		expectedErr error
	}{
		{
			name:        "date is invalid format",
			fromDate:    "january 10th, 2020",
			toDate:      "february 10th, 2021",
			doMocks:     func() {},
			expected:    nil,
			expectedErr: service.ErrInvalidDateTimeFormat,
		},
		{
			name:     "success get balance histories",
			fromDate: "2020-01-13T14:00:00+07:00",
			toDate:   "2021-01-13T14:00:00+07:00",
			doMocks: func() {
				suite.
					mockTransactionRepository.
					EXPECT().
					GetBalanceHistory(gomock.Any(), "2020-01-13T14:00:00+07:00", "2021-01-13T14:00:00+07:00").
					Return([]model.BalanceHistory{{}, {}}, nil).
					Times(1)
			},
			expected:    []model.BalanceHistory{{}, {}},
			expectedErr: nil,
		},
		{
			name:     "failed get balance histories",
			fromDate: "2020-01-13T14:00:00+07:00",
			toDate:   "2021-01-13T14:00:00+07:00",
			doMocks: func() {
				suite.
					mockTransactionRepository.
					EXPECT().
					GetBalanceHistory(gomock.Any(), "2020-01-13T14:00:00+07:00", "2021-01-13T14:00:00+07:00").
					Return(nil, sql.ErrConnDone).
					Times(1)
			},
			expected:    nil,
			expectedErr: sql.ErrConnDone,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tt.doMocks()

			service := service.NewTransactionService(suite.mockTransactionRepository)
			actual, actualErr := service.GetBalanceHistory(ctx, tt.fromDate, tt.toDate)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}

func (suite *transactionTestsSuite) SetupTest() {
	t := suite.T()
	suite.ctrl = gomock.NewController(t)
	suite.mockTransactionRepository = mock.NewMockTransactionRepository(suite.ctrl)
}

func (suite *transactionTestsSuite) TearDownTest() {
	defer suite.ctrl.Finish()
}
