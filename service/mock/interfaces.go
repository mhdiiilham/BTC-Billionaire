// Code generated by MockGen. DO NOT EDIT.
// Source: service/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mhdiiilham/BTC-Billionaire/model"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// RecordTransaction mocks base method.
func (m *MockTransactionRepository) RecordTransaction(ctx context.Context, transaction model.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordTransaction", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordTransaction indicates an expected call of RecordTransaction.
func (mr *MockTransactionRepositoryMockRecorder) RecordTransaction(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).RecordTransaction), ctx, transaction)
}