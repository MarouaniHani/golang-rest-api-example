package sql

import (
	"context"
	"errors"
	"fmt"
	"kuwait-test/dto"
	"kuwait-test/model"
)

func (c *conn) IsOrderPaymentExist(ctx context.Context, orderId string) (bool, error) {
	if len(orderId) == 0 {
		return false, errors.New("please enter and order id")
	}
	query := `SELECT count(*) FROM order_payments WHERE order_id = ?`
	var count int
	if err := c.db.QueryRow(query, orderId).Scan(&count); err != nil {
		return false, err
	}
	if count >0 {
		return true, nil
	}
	return false, nil
}
func (c *conn) GetOrderPayment(ctx context.Context, orderId string) (*model.OrderPayment, error){
	if len(orderId) == 0 {
		return nil, errors.New("please enter and order id")
	}
	query := `SELECT * FROM order_payments WHERE order_id = ?`
	row := c.db.QueryRow(query, orderId)
	orderPayment := model.OrderPayment{}
	err := row.Scan(&orderPayment.OrderID, &orderPayment.Sequential,&orderPayment.Method, &orderPayment.Installments, &orderPayment.Amount)
	if err != nil {
		return nil, err
	}
	return &orderPayment, nil
}

func (c *conn) PayOrder(ctx context.Context, payOrder dto.PayOrder, orderId string) (string, error) {
	if len(payOrder.Method) == 0 {
		return "", errors.New("please enter payment method")
	}
	if payOrder.Amount <= 0 {
		return "", errors.New("please enter a correct payment amount")
	}
	if len(orderId) == 0 {
		return "", errors.New("please enter and order id")
	}
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
	query := `INSERT INTO order_payments(order_id, sequential, method, installments, amount) VALUES (?, ?, ?, ?,?)`
	if res, err := tx.Exec(query, orderId, 1, payOrder.Method,0,payOrder.Amount); err != nil {
		return "", err
	} else {
		fmt.Println("res", res)
	}
	return orderId, nil
}
