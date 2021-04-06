package models

import "context"

// Order represent orders data
type Order struct {
	OrderID     uint32  `json:"order_id"`
	CustomerID  uint32  `json:"customer_id"`
	OrderDate   string  `json:"order_date"`
	TotalAmount float32 `json:"total_amount"`
	OrderStatus int32   `json:"order_status"` // 0:pending, 1:confirmed, 2:cancelled
	ProcessedBy uint32  `json:"processed_by"`
}

// OrderDetails represent order_details data
type OrderDetails struct {
	OrderDetailsID uint32  `json:"order_deatils_id"`
	OrderID        uint32  `json:"order_id"`
	MenuID         uint32  `json:"menu_id"`
	AmountID       float32 `json:"amount"`
	NumOfServing   int32   `json:"no_of_serving"`
	TotalAmount    uint32  `json:"total_amount"`
}

// OrderRepository represents repo object for order
type OrderRepository interface {
	GetAll() ([]Order, error)
	GetById(OrderID uint32) (*Order, error)
	Update(ctx context.Context, order *Order) error
	DeleteById(ctx context.Context, OrderID uint32) error
	Store(ctx context.Context, ord *Order) error
}

// OrderDetailsRepository represents repo object for order
type OrderDetailsRepository interface {
	GetAll() ([]OrderDetails, error)
	GetById(OrderID uint32) (*OrderDetails, error)
	Update(OrderID uint32, order *OrderDetails) error
	DeleteById(OrderID uint32) error
	Store(ctx context.Context, ord *Order) error
}
