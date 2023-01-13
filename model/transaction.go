package model

import "time"

type Transaction struct {
	ID       int       `db:"id" json:"-"`
	Datetime time.Time `db:"datetime" json:"datetime"`
	Amount   float64   `db:"amount" json:"amount"`
}

type TransactionRequest struct {
	Datetime string  `json:"datetime"`
	Amount   float64 `json:"amount"`
}
