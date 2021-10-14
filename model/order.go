package model

import "time"

type Order struct {
	ID                string    `json:"id" json:"id"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	Status            string    `json:"status" db:"status"`
	PurchasedAt       time.Time `json:"purchased_at" db:"purchased_at"`
	ApprovedAt        time.Time `json:"approved_at" db:"approved_at"`
	PickeupAt         time.Time `json:"pickeup_at" db:"pickeup_at"`
	DeliveredAt       time.Time `json:"delivered_at" db:"delivered_at"`
	EstimatedDelivery time.Time `json:"estimated_delivery" db:"estimated_delivery"`
}
