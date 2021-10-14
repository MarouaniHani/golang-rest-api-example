package sql

import (
	"context"
	"errors"
	"fmt"
	"kuwait-test/model"
	"strings"
)

func (c *conn) ListProductsByIDs(ctx context.Context, ids []string) ([]*model.Product, error) {

	if len(ids) == 0 {
		return nil, errors.New("please enter at least one product id")
	}
	d := strings.Join(ids, "','")
	query := fmt.Sprintf(`SELECT * FROM products WHERE id IN ('%s')`, d)
	rows, err := c.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var products []*model.Product

	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.ID, &product.Category, &product.NameLen, &product.DescLen, &product.Photos, &product.Weight, &product.Length, &product.Height, &product.Width)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
