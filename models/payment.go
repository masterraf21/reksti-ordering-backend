package models

import (
	"context"
)

// Payment huhu
type Payment struct {
	PaymentID     uint32  `json:"payment_id"`
	OrderID       uint32  `json:"order_id"`
	Amount        float32 `json:"amount"`
	PaymentTypeID uint32  `json:"payment_type_id"`
	PaymentDate   string  `json:"payment_date"`
	PaymentStatus int32   `json:"payment_status"` // 0:pending,1:completed,1:cancelled
}

// PaymentRepository pymnerepo
type PaymentRepository interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(ctx context.Context, paymentID uint32) error
	Store(ctx context.Context, payment *Payment) (uint32, error)
	UpdateStatus(ctx context.Context, paymentID uint32, status int32) error
	GetListOfPaymentsByCustomerID(ctx context.Context, customerID uint32) ([]Payment, error)
}

// PaymentUsecase for payment
type PaymentUsecase interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(ctx context.Context, paymentID uint32) error
	CreatePayment(ctx context.Context, payment *Payment) (uint32, error)
	UpdateStatus(ctx context.Context, paymentID uint32, status int32) error
	GetListOfPaymentsByCustomerID(ctx context.Context, customerID uint32) ([]Payment, error)
}
