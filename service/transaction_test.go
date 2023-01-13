package service_test

import (
	"context"
	"database/sql"
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
		doMocks  func()
		expected error
	}{
		{
			name:     "failed to covert dateTime string to time.Time",
			dateTime: "january the third",
			amount:   10,
			doMocks:  func() {},
			expected: service.ErrInvalidDateTimeFormat,
		},
		{
			name:     "success record new transaction to database",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   10,
			doMocks: func() {

				trxTime, _ := time.Parse(time.RFC3339, "2019-10-05T14:45:05+07:00")
				suite.
					mockTransactionRepository.
					EXPECT().
					RecordTransaction(gomock.Any(), model.Transaction{
						Datetime: trxTime,
						Amount:   10,
					}).Return(nil).
					Times(1)
			},
			expected: nil,
		},
		{
			name:     "failed record new transaction to database",
			dateTime: "2019-10-05T14:45:05+07:00",
			amount:   10,
			doMocks: func() {

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
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tt.doMocks()

			svc := service.NewTransactionService(suite.mockTransactionRepository)
			actual := svc.RecordNewTransaction(ctx, tt.dateTime, tt.amount)
			assert.Equal(t, tt.expected, actual, "expected function NewTransactionRecord returning %v", tt.expected)
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
