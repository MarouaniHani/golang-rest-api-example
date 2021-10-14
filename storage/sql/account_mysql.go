package sql

import (
	"context"
	"errors"
	"fmt"
	"kuwait-test/model"
)

func(c *conn) SaveAccount(ctx context.Context, data *model.Account)(string,error){
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
	if data == nil {
		return "", errors.New("data is empty")
	}
	insertAccountSQL := `INSERT INTO accounts(customer_id,username,password) VALUES (?, ?, ?) `
	if res,err := tx.Exec(insertAccountSQL,data.CustomerID,data.UserName,data.Password); err != nil {
		return "", err
	}else {
		fmt.Println("res",res)
	}

	return data.CustomerID, nil
}
func(c *conn) IsCustomerIdExist(ctx context.Context, customerId string) (bool, error) {
	var count int
	if err := c.db.QueryRow("SELECT count(*) FROM customers WHERE id = ?", customerId).Scan(&count); err != nil {
		return false, err
	}
	if count >0 {
		return true, nil
	}
	return false, nil
}

func(c *conn) IsRegistered(ctx context.Context, customerId string) (bool, error) {
	var count int
	if err := c.db.QueryRow("SELECT count(*) FROM accounts WHERE customer_id = ?", customerId).Scan(&count); err != nil {
		return false, err
	}
	if count >0 {
		return false, nil
	}
	return true, nil
}

func(c *conn) GetAccountByID(ctx context.Context, customerId string) (*model.Account, error) {
	var account model.Account
	if err := c.db.QueryRow("SELECT * FROM accounts WHERE customer_id = ?", customerId).Scan(&account.CustomerID,&account.UserName,&account.Password); err != nil {
		return nil, err
	}
	return &account, nil
}