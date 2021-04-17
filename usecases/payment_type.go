package usecases

import (
	"context"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type paymentTypeUsecases struct {
	PaymentTypeRepo models.PaymentTypeRepository
}

// NewPaymentTypeUsecase will initiate usecase
func NewPaymentTypeUsecase(ptr models.PaymentTypeRepository) models.PaymentTypeUsecase {
	return &paymentTypeUsecases{
		PaymentTypeRepo: ptr,
	}
}

func (u *paymentTypeUsecases) GetAll() (res []models.PaymentType, err error) {
	res, err = u.PaymentTypeRepo.GetAll()
	return
}

func (u *paymentTypeUsecases) GetByID(paymentTypeID uint32) (res *models.PaymentType, err error) {
	res, err = u.PaymentTypeRepo.GetByID(paymentTypeID)
	return
}

func (u *paymentTypeUsecases) DeleteByID(ctx context.Context, paymentTypeID uint32) error {
	err := u.PaymentTypeRepo.DeleteByID(ctx, paymentTypeID)
	return err
}

func (u *paymentTypeUsecases) CreatePaymentType(
	ctx context.Context,
	paymentType *models.PaymentType,
) (res uint32, err error) {
	res, err = u.PaymentTypeRepo.Store(ctx, paymentType)
	return
}
