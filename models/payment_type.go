package models

import "context"

// PaymentType for payment type
type PaymentType struct {
	PaymentTypeID uint32 `json:"payment_type_id"`
	Method        string `json:"payment_method"`
	Company       string `json:"payment_company"`
}

// PaymentTypeRepository for repo
type PaymentTypeRepository interface {
	GetAll() ([]PaymentType, error)
	GetByID(paymentTypeID uint32) (*PaymentType, error)
	DeleteByID(ctx context.Context, paymentTypeID uint32) error
	Store(ctx context.Context, paymentType *PaymentType) (uint32, error)
}

// PaymentTypeUsecase for usecase
type PaymentTypeUsecase interface {
	GetAll() ([]PaymentType, error)
	GetByID(paymentTypeID uint32) (*PaymentType, error)
	DeleteByID(ctx context.Context, paymentTypeID uint32) error
	CreatePaymentType(ctx context.Context, paymentType *PaymentType) (uint32, error)
}
