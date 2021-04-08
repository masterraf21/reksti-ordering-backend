package models

import "context"

// Order represent orders data
type Order struct {
	OrderID     uint32  `json:"order_id"`
	CustomerID  uint32  `json:"customer_id"`
	OrderDate   string  `json:"order_date"`
	TotalPrice  float32 `json:"total_price"`
	OrderStatus int32   `json:"order_status,omitempty"` // 0:pending, 1:completed, 2:cancelled
}

// OrderDetails represent order_details data
type OrderDetails struct {
	OrderDetailsID uint32  `json:"order_deatils_id"`
	OrderID        uint32  `json:"order_id"`
	MenuID         uint32  `json:"menu_id"`
	Quantity       uint32  `json:"quantity"`
	TotalPrice     float32 `json:"total_price"`
}

// OrderWithDetails for aggregating
type OrderWithDetails struct {
	OrderInfo Order          `json:"order_info"`
	Details   []OrderDetails `json:"details"`
}

// OrderRepository represents repo object for order
type OrderRepository interface {
	GetByStatusAndCustID(status int32, custID uint32) ([]Order, error)
	GetByCustID(custID uint32) ([]Order, error)
	GetAll() ([]Order, error)
	GetByID(OrderID uint32) (*Order, error)
	UpdateByID(ctx context.Context, orderID uint32, order *Order) error
	UpdateOrderStatus(ctx context.Context, orderID uint32, status int32) error
	UpdateTotalPrice(ctx context.Context, orderID uint32) error
	DeleteByID(ctx context.Context, OrderID uint32) error
	Store(ctx context.Context, ord *Order) (uint32, error)
	BulkInsert(ctx context.Context, orders []Order) error
}

// OrderDetailsRepository represents repo object for order
type OrderDetailsRepository interface {
	UpdateTotalPrice(orderDetailsID uint32) error
	GetOrderDetailsByOrderID(orderID uint32) []OrderDetails
	GetAll() ([]OrderDetails, error)
	GetByID(OrderID uint32) (*OrderDetails, error)
	Update(OrderID uint32, order *OrderDetails) error
	DeleteByID(OrderID uint32) error
	Store(ctx context.Context, ord *Order) error
	BulkInsert(ctx context.Context, orderdetails []OrderDetails) error
}

// OrderUsecase to create usecase for order
type OrderUsecase interface {
	GetOrderWithDetails(orderID uint32) (OrderWithDetails, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID uint32) (*Order, error)
	GetOrdersByCustID(customerID uint32) ([]Order, error)
	GetOrdersHistoryByCustID(customerID uint32) ([]Order, error)
	GetOngoingOrdersyByCustID(customerID uint32) ([]Order, error)
	CreateOrder(ctx context.Context, order *Order) error
	BulkCreateOrders(ctx context.Context, orders []Order) error
	UpdateOrder(ctx context.Context, orderID, order *Order) error
}
