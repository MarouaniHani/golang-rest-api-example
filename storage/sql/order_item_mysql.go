package sql

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (c *conn) SaveOrderItem(ctx context.Context, orderID string, productID string) (string, error) {
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
	shippingDate := time.Now()
	insertOrderItemSQL := `INSERT INTO order_items(order_id, order_item_id, product_id, seller_id, shipping_date, price, freight) 
							VALUES (?, ?, ?, '0', ?, 0, 0) `
	if res, err := tx.Exec(insertOrderItemSQL, orderID, strconv.Itoa(generatedId), productID,shippingDate); err != nil {
		return "", err
	} else {
		fmt.Println("res", res)
	}

	return strconv.Itoa(generatedId), nil
}
