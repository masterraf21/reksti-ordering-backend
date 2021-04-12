package models

import "context"

// Order represent orders data
type Order struct {
	OrderID     uint32  `json:"order_id"` // auto increment
	CustomerID  uint32  `json:"customer_id"`
	OrderDate   string  `json:"order_date"`
	TotalPrice  float32 `json:"total_price"`
	OrderStatus int32   `json:"order_status,omitempty"` // 0:pending, 1:completed, 2:cancelled
}

// OrderDetails represent order_details data
type OrderDetails struct {
	OrderDetailsID uint32  `json:"order_details_id"`
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
	GetOrderDetailsByOrderID(orderID uint32) ([]OrderDetails, error)
	GetAll() ([]Order, error)
	GetByID(OrderID uint32) (*Order, error)
	UpdateArbitrary(
		ctx context.Context,
		orderID uint32,
		columnName string,
		value interface{},
	) error
	UpdateByID(ctx context.Context, orderID uint32, order *Order) error
	UpdateOrderStatus(ctx context.Context, orderID uint32, status int32) error
	UpdateTotalPrice(ctx context.Context, orderID uint32) error
	DeleteByID(ctx context.Context, OrderID uint32) error
	Store(ctx context.Context, ord *Order) (uint32, error)
	BulkInsert(ctx context.Context, orders []Order) error
}

// OrderDetailsRepository represents repo object for order
type OrderDetailsRepository interface {
	GetOrderDetailsByOrderID(orderID uint32) ([]OrderDetails, error)
	GetAll() ([]OrderDetails, error)
	GetByID(orderDetailsID uint32) (*OrderDetails, error)
	UpdateByID(ctx context.Context, orderDetailsID uint32, order *OrderDetails) error
	UpdateTotalPrice(ctx context.Context, orderDetailsID uint32) error
	UpdateArbitrary(
		ctx context.Context,
		orderDetailsID uint32,
		columnName string,
		value interface{},
	) error
	DeleteByID(ctx context.Context, orderDetailsID uint32) error
	Store(ctx context.Context, ord *OrderDetails) (uint32, error)
	BulkInsert(ctx context.Context, orderdetails []OrderDetails) error
}

// OrderUsecase to create usecase for order
type OrderUsecase interface {
	GetOrderWithDetails(orderID uint32) (OrderWithDetails, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID uint32) (*Order, error)
	GetOrderDetailByID(orderID uint32) (*OrderDetails, error)
	GetOrdersByCustID(customerID uint32) ([]Order, error)
	GetOrdersHistoryByCustID(customerID uint32) ([]Order, error)
	GetOngoingOrdersyByCustID(customerID uint32) ([]Order, error)
	CreateOrder(ctx context.Context, order *Order) (uint32, error)
	BulkCreateOrders(ctx context.Context, orders []Order) error
	UpdateOrder(ctx context.Context, orderID uint32, order *Order) error
	UpdateOrderPrice(ctx context.Context, orderID uint32) error
	UpdateOrderDetailPrice(ctx context.Context, orderDetailID uint32) error
	UpdateOrderStatus(ctx context.Context, orderID uint32, status int32) error
	CreateOrderDetail(ctx context.Context, orderD *OrderDetails) (uint32, error)
	DeleteOrder(ctx context.Context, orderID uint32) error
	DeleteOrderDetail(ctx context.Context, orderDetailID uint32) error
}
