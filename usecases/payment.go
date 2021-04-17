package usecases

import (
	"context"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type paymentUsecase struct {
	PaymentRepo models.PaymentRepository
}

// NewPaymentUsecase will initiate new usecase
func NewPaymentUsecase(pyr models.PaymentRepository) models.PaymentUsecase {
	return &paymentUsecase{
		PaymentRepo: pyr,
	}
}

func (u *paymentUsecase) UpdateStatus(ctx context.Context, paymentID uint32, status int32) error {
	err := u.PaymentRepo.UpdateStatus(ctx, paymentID, status)
	return err
}

func (u *paymentUsecase) GetAll() (res []models.Payment, err error) {
	res, err = u.PaymentRepo.GetAll()
	return
}

func (u *paymentUsecase) GetByID(paymentID uint32) (res *models.Payment, err error) {
	res, err = u.PaymentRepo.GetByID(paymentID)
	return
}

func (u *paymentUsecase) DeleteByID(ctx context.Context, paymentID uint32) error {
	err := u.PaymentRepo.DeleteByID(ctx, paymentID)
	return err
}

func (u *paymentUsecase) CreatePayment(ctx context.Context, payment *models.Payment) (id uint32, err error) {
	id, err = u.PaymentRepo.Store(ctx, payment)
	return
}
