package rest

import "context"

type Transactioner interface {
	RecordNewTransaction(ctx context.Context, dateTime string, amount float64) error
}
