package pkg

import (
	"errors"
	"kuwait-test/model"
)

func ValidAccount(account model.Account) error {
	if len(account.UserName) == 0 {
		return errors.New("you should enter user name")
	}
	if len(account.Password) == 0 {
		return errors.New("you should enter password")
	}
	return nil
}
