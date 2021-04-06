package models

import "context"

// Customer for customer table
type Customer struct {
	CustomerID    uint32 `json:"customer_id"`
	FullName      string `json:"customer_full_name"`
	Email         string `json:"customer_email"`
	PhoneNumber   string `json:"customer_phone_numer"`
	Username      string `json:"customer_username"`
	Password      string `json:"password"`
	AccountStatus bool   `json:"account_status"`
}

// CustomerRepository for repo
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(CustomerID uint32) (*Customer, error)
	UpdateById(ctx context.Context, CustomerID uint32) error
	DeleteById(ctx context.Context, CustomerID uint32) error
	Store(ctx context.Context, cust *Customer) error
}
