package sql

import (
	"context"
	"errors"
	"fmt"
	"kuwait-test/model"
	"math/rand"
	"strconv"
	"time"
)

func (c *conn) ListOrdersByCustomerID(ctx context.Context, id string) ([]*model.Order, error) {

	if len(id) == 0 {
		return nil, errors.New("customer id is empty")
	}
	query := `SELECT * FROM orders WHERE customer_id =?`

	rows, err := c.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order

	for rows.Next() {
		order := model.Order{}
		err = rows.Scan(&order.ID, &order.CustomerID, &order.Status, &order.PurchasedAt, &order.ApprovedAt, &order.PickeupAt, &order.DeliveredAt, &order.EstimatedDelivery)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (c *conn) ListOrderItemsByOrderID(ctx context.Context, orderId string) ([]*model.OrderItem, error) {

	if len(orderId) == 0 {
		return nil, errors.New("order id is empty")
	}
	query := `SELECT * FROM order_items WHERE order_id =?`

	rows, err := c.db.Query(query, orderId)
	if err != nil {
		return nil, err
	}

	var items []*model.OrderItem

	for rows.Next() {
		item := model.OrderItem{}
		err = rows.Scan(&item.OrderID, &item.OrderItemID, &item.ProductID, &item.SellerID, &item.ShippingDate, &item.Price, &item.Freight)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (c *conn) ListOrderPaymentsByOrderID(ctx context.Context, orderId string) ([]*model.OrderPayment, error) {

	if len(orderId) == 0 {
		return nil, errors.New("order id is empty")
	}
	query := `SELECT * FROM order_payments WHERE order_id =?`

	rows, err := c.db.Query(query, orderId)
	if err != nil {
		return nil, err
	}

	var items []*model.OrderPayment

	for rows.Next() {
		item := model.OrderPayment{}
		err = rows.Scan(&item.OrderID, &item.Sequential, &item.Method, &item.Installments, &item.Amount)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (c *conn) GetOrderByCustomerIDAndOrderID(ctx context.Context, customerId string, orderId string) (*model.Order, error) {

	if len(customerId) == 0 {
		return nil, errors.New("customer id is empty")
	}

	if len(orderId) == 0 {
		return nil, errors.New("order id is empty")
	}

	query := `SELECT * FROM orders WHERE customer_id = ? and id =?`

	rows := c.db.QueryRow(query, customerId, orderId)

		order := model.Order{}
		err := rows.Scan(&order.ID, &order.CustomerID, &order.Status, &order.PurchasedAt, &order.ApprovedAt, &order.PickeupAt, &order.DeliveredAt, &order.EstimatedDelivery)
		if err != nil {
			return nil, err
		}
	return &order, nil
}

func (c *conn) SaveOrder(ctx context.Context,customerID string) (string, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	generatedId := rand.New(
		rand.NewSource(time.Now().UnixNano())).Int()
	insertAccountSQL := `INSERT INTO orders(id,customer_id,status, purchased_at, approved_at, pickedup_at, delivered_at, estimated_delivery) 
							VALUES (?, ?, 'submitted','2017-08-01 19:00:07', '2017-08-01 19:00:07', '2017-08-01 19:00:07', '2017-08-01 19:00:07', '2017-08-01 19:00:07') `
	if res, err := tx.Exec(insertAccountSQL,strconv.Itoa(generatedId), customerID); err != nil {
		return "", err
	} else {
		fmt.Println("res", res)
	}

	return strconv.Itoa(generatedId), nil
}
