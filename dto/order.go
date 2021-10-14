package dto

import (
	"time"
)

type Order struct {
	ID          string         `json:"id" json:"id"`
	Status      string         `json:"status" db:"status"`
	PurchasedAt time.Time      `json:"purchased_at" db:"purchased_at"`
	PickedupAt  time.Time      `json:"pickedup_at" db:"pickedup_at"`
	DeliveredAt time.Time      `json:"delivered_at" db:"delivered_at"`
	Items       []OrderItem    `json:"items"`
	Payment     OrderPayment `json:"payment"`
}
