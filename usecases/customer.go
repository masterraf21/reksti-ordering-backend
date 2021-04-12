package usecases

import (
	"context"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type customerUsecase struct {
	CustomerRepo models.CustomerRepository
}

// NewCustomerUsecase will initiate new usecase
func NewCustomerUsecase(cur models.CustomerRepository) models.CustomerUsecase {
	return &customerUsecase{
		CustomerRepo: cur,
	}
}

func (u *customerUsecase) GetAll() (res []models.Customer, err error) {
	res, err = u.CustomerRepo.GetAll()
	return
}

func (u *customerUsecase) GetByID(CustomerID uint32) (res *models.Customer, err error) {
	res, err = u.CustomerRepo.GetByID(CustomerID)
	return
}

func (u *customerUsecase) DeleteByID(ctx context.Context, CustomerID uint32) error {
	err := u.CustomerRepo.DeleteByID(ctx, CustomerID)
	return err
}

func (u *customerUsecase) CreateCustomer(ctx context.Context, cust *models.Customer) (id uint32, err error) {
	id, err = u.CustomerRepo.Store(ctx, cust)
	return
}
