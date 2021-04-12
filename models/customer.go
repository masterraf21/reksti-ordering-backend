package models

import "context"

// Customer for customer table
type Customer struct {
	CustomerID    uint32 `json:"customer_id"`
	FullName      string `json:"customer_full_name"`
	Email         string `json:"customer_email"`
	PhoneNumber   string `json:"customer_phone_numer"`
	Username      string `json:"customer_username"`
	Password      string `json:"customer_password"`
	AccountStatus bool   `json:"account_status"`
}

// CustomerRepository for repo
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(CustomerID uint32) (*Customer, error)
	DeleteByID(ctx context.Context, CustomerID uint32) error
	Store(ctx context.Context, cust *Customer) (uint32, error)
}

// CustomerUsecase for usecase
type CustomerUsecase interface {
	GetAll() ([]Customer, error)
	GetByID(CustomerID uint32) (*Customer, error)
	DeleteByID(ctx context.Context, CustomerID uint32) error
	CreateCustomer(ctx context.Context, cust *Customer) (uint32, error)
}
