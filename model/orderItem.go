package model

import "time"

type OrderItem struct {
	OrderID      string    `json:"order_id" db:"order_id"`
	OrderItemID  string    `json:"order_item_id" db:"order_item_id"`
	ProductID    string    `json:"product_id" db:"product_id"`
	SellerID     string    `json:"seller_id" db:"seller_id"`
	ShippingDate time.Time `json:"shipping_date" db:"shipping_date"`
	Price        float64   `json:"price" db:"price"`
	Freight      float64   `json:"freight" db:"freight"`
}
