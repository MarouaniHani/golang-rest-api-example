package dto

type PayOrder struct {
	Method       string  `json:"method" db:"method"`
	Amount       float64 `json:"amount" db:"amount"`
}
