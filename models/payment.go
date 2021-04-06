package models

import "context"

// Payment huhu
type Payment struct {
	PaymentID   uint32  `json:"payment_id"`
	OrderID     uint32  `json:"order_id"`
	Amount      float32 `json:"amount"`
	PaidBy      string  `json:"paid_by"`
	PaymentDate string  `json:"payment_date"`
	ProcessedBy uint32  `json:"processed_by"`
}

// PaymentRepository pymnerepo
type PaymentRepository interface {
	GetAll() ([]Payment, error)
	GetById(paymentID uint32) (*Payment, error)
	Update(ctx context.Context, payment *Payment) error
	DeleteById(paymentID uint32) error
	Store(payment *Payment) error
}
