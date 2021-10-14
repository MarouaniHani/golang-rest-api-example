package utils

import (
	"kuwait-test/dto"
	"kuwait-test/model"
	"math/rand"
)

func MapOrderDto(order *model.Order, items []*model.OrderItem, payment *model.OrderPayment) dto.Order {
	var itemsDto []dto.OrderItem
	if items != nil {
		for _, item := range items {
			itemsDto = append(itemsDto, dto.OrderItem{
				Price:   item.Price,
				Freight: item.Freight,
			})
		}

	}
	var paymentDto dto.OrderPayment
	if payment != nil {
		paymentDto = dto.OrderPayment{
			Method: payment.Method,
			Amount: payment.Amount,
		}
	}

	return dto.Order{
		ID:          order.ID,
		Status:      order.Status,
		PurchasedAt: order.PurchasedAt,
		PickedupAt:  order.PickeupAt,
		DeliveredAt: order.DeliveredAt,
		Items:       itemsDto,
		Payment:     paymentDto,
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
