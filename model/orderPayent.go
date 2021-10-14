package model

type OrderPayment struct {
	OrderID      string  `json:"order_id" db:"order_id"`
	Sequential   int     `json:"sequential" db:"sequential"`
	Method       string  `json:"method" db:"method"`
	Installments int     `json:"installments" db:"installments"`
	Amount       float64 `json:"amount" db:"amount"`
}
