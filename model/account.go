package model

type Account struct {
	CustomerID string `json:"customer_id,omitempty" db:"customer_id"`
	UserName   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
}
