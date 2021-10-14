package dto

type OrderPayment struct {
	Method       string  `json:"method"`
	Amount       float64 `json:"amount"`
}
