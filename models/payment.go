package models

import "context"

// Payment huhu
type Payment struct {
	PaymentID   uint32  `json:"payment_id"`
	OrderID     uint32  `json:"order_id"`
	Amount      float32 `json:"amount"`
	PaymentType string  `json:"paymeny_type"`
	PaymentDate string  `json:"payment_date"`
}

// PaymentRepository pymnerepo
type PaymentRepository interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(ctx context.Context, paymentID uint32) error
	Store(ctx context.Context, payment *Payment) error
}

// PaymentUsecase for payment
type PaymentUsecase interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(ctx context.Context, paymentID uint32) error
	CreatePayment(ctx context.Context, payment *Payment) error
}
