package rest

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"kuwait-test/model"
	"kuwait-test/pkg"
)

func (s *Server) CreateAccount(c *gin.Context) {
	var req model.Account
	if err := c.ShouldBind(&req); err != nil {
		BindJsonErr(c, err)
		return
	}
	if err := pkg.ValidAccount(req); err != nil {
		BindJsonErr(c, err)
		return
	}

	exist, err := s.storage.IsCustomerIdExist(c, req.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		AbortWithError(c, 500, `isCustomerIdExist failed`, err)
		return
	}

	if exist == false {
		BindJsonErr(c, errors.New("please enter a correct customer id"))
		return
	}

	isRegistered, err := s.storage.IsRegistered(c, req.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		AbortWithError(c, 500, `isRegistered failed`, err)
		return
	}

	if isRegistered == false {
		BindJsonErr(c, errors.New("account already registered"))
		return
	}

	customerID, err := s.storage.SaveAccount(c, &req)
	if err != nil {
		AbortWithError(c, 500, "failed to save account ", err)
		return
	}
	if customerID == "" {
		ResponseData(c, nil)
	}
	account, err := s.storage.GetAccountByID(c, customerID)
	if err != nil {
		AbortWithError(c, 500, "failed to get account by id", err)
		return
	}

	ResponseData(c, account)
}