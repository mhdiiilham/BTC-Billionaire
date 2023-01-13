package model

import "time"

type BalanceHistory struct {
	Datetime time.Time `db:"datetime" json:"datetime"`
	Amount   float64   `db:"amount" json:"amount"`
}
