package storage

import (
	"context"
	"github.com/pkg/errors"
	"kuwait-test/dto"
	"kuwait-test/model"
)

var (
	// ErrNotFound is the error returned by storages if a resource cannot be found.
	ErrNotFound = errors.New("not found")
)

// Storage is the storage interface used by the server. Implementations are
// required to be able to perform atomic compare-and-swap updates and support standardize on UTC.
type Storage interface {
	Close() error
	Version() (string, error)

	SaveAccount(ctx context.Context, account *model.Account) (string, error)
	IsCustomerIdExist(ctx context.Context, customerId string) (bool, error)
	IsRegistered(ctx context.Context, customerId string) (bool, error)
	GetAccountByID(ctx context.Context, customerId string) (*model.Account, error)

	ListOrdersByCustomerID(ctx context.Context, customerId string) ([]*model.Order, error)
	GetOrderByCustomerIDAndOrderID(ctx context.Context, customerId string, orderId string) (*model.Order, error)
	ListOrderItemsByOrderID(ctx context.Context, orderId string) ([]*model.OrderItem, error)
	SaveOrder(ctx context.Context, customerID string) (string, error)
	SaveOrderItem(ctx context.Context, orderID string, productID string) (string, error)

	ListProductsByIDs(ctx context.Context, ids []string) ([]*model.Product, error)
	IsOrderPaymentExist(ctx context.Context, orderId string) (bool, error)
	GetOrderPayment(ctx context.Context, orderId string) (*model.OrderPayment, error)
	PayOrder(ctx context.Context, payOrder dto.PayOrder, orderId string) (string, error)
}
