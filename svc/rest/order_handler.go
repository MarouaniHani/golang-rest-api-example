package rest

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"kuwait-test/dto"
	"kuwait-test/utils"
)

func (s *Server) ListOrdersByCustomerID(c *gin.Context) {
	customerID := c.Param(`id`)
	if len(customerID) == 0 {
		BindJsonErr(c, errors.New("customer id should not be empty"))
		return
	}
	data, err := s.storage.ListOrdersByCustomerID(c, customerID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders by customer id`, err)
		return
	}
	var orders []dto.Order
	for _, order := range data {
		orderItems, err := s.storage.ListOrderItemsByOrderID(c, order.ID)
		if err != nil {
			AbortWithError(c, 500, `unable to get list of orders items by order id`, err)
			return
		}
			orderPayment, _ := s.storage.GetOrderPayment(c, order.ID)

		orders = append(orders, utils.MapOrderDto(order, orderItems, orderPayment))
	}
	ResponseData(c, orders)
}

func (s *Server) GetOrderByCustomerID(c *gin.Context) {
	customerID := c.Param(`customer_id`)
	if len(customerID) == 0 {
		BindJsonErr(c, errors.New("customer id should not be empty"))
		return
	}
	orderID := c.Param(`order_id`)
	if len(customerID) == 0 {
		BindJsonErr(c, errors.New("order id should not be empty"))
		return
	}
	data, err := s.storage.GetOrderByCustomerIDAndOrderID(c, customerID, orderID)
	if err != nil {
		AbortWithError(c, 500, `unable to get order by customer id`, err)
		return
	}
	orderItems, err := s.storage.ListOrderItemsByOrderID(c, data.ID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders items by order id`, err)
		return
	}
	orderPayment, err := s.storage.GetOrderPayment(c, data.ID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders payments by order id`, err)
		return
	}

	ResponseData(c, utils.MapOrderDto(data, orderItems, orderPayment))
}

func (s *Server) CreateOrder(c *gin.Context) {
	customerID := c.Param(`customer_id`)
	if len(customerID) == 0 {
		BindJsonErr(c, errors.New("customer id should not be empty"))
		return
	}
	var req dto.CreateOrder
	if err := c.ShouldBind(&req); err != nil {
		BindJsonErr(c, err)
		return
	}

	products, err := s.storage.ListProductsByIDs(c, req.ProductIDs)
	if err != nil && err != sql.ErrNoRows {
		AbortWithError(c, 500, `ListProductsByIDs failed`, err)
		return
	}
	exist, err := s.storage.IsCustomerIdExist(c, customerID)
	if err != nil && err != sql.ErrNoRows {
		AbortWithError(c, 500, `isCustomerIdExist failed`, err)
		return
	}

	if exist == false {
		BindJsonErr(c, errors.New("please enter a correct customer id"))
		return
	}

	orderID, err := s.storage.SaveOrder(c, customerID)
	if err != nil {
		AbortWithError(c, 500, "failed to save order", err)
		return
	}

	for _, product := range products {
		_, err = s.storage.SaveOrderItem(c, orderID, product.ID)
		if err != nil {
			AbortWithError(c, 500, "failed to save order item", err)
			return
		}
	}

	data, err := s.storage.GetOrderByCustomerIDAndOrderID(c, customerID, orderID)
	if err != nil {
		AbortWithError(c, 500, "failed to get order by id and customer id", err)
		return
	}
	orderItems, err := s.storage.ListOrderItemsByOrderID(c, data.ID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders items by order id`, err)
		return
	}

	ResponseData(c, utils.MapOrderDto(data, orderItems, nil))
}

func (s *Server) PayOrder(c *gin.Context) {
	customerID := c.Param(`customer_id`)
	if len(customerID) == 0 {
		BindJsonErr(c, errors.New("customer id should not be empty"))
		return
	}
	orderID := c.Param(`order_id`)
	if len(orderID) == 0 {
		BindJsonErr(c, errors.New("order id should not be empty"))
		return
	}
	var req dto.PayOrder
	if err := c.ShouldBind(&req); err != nil {
		BindJsonErr(c, err)
		return
	}
	exist, err := s.storage.IsOrderPaymentExist(c, orderID)
	if err != nil {
		AbortWithError(c, 500, "failed to get order payment", err)
		return
	}
	if exist == true {
		ResponseData(c, "order already paid")
		return
	}
	_, err = s.storage.PayOrder(c, req, orderID)
	if err != nil {
		AbortWithError(c, 500, "failed to pay order", err)
		return
	}
	data, err := s.storage.GetOrderByCustomerIDAndOrderID(c, customerID, orderID)
	if err != nil {
		AbortWithError(c, 500, "failed to get order by id and customer id", err)
		return
	}
	orderItems, err := s.storage.ListOrderItemsByOrderID(c, data.ID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders items by order id`, err)
		return
	}
	orderPayment, err := s.storage.GetOrderPayment(c, orderID)
	if err != nil {
		AbortWithError(c, 500, `unable to get list of orders items by order id`, err)
		return
	}

	ResponseData(c, utils.MapOrderDto(data, orderItems, orderPayment))
}
