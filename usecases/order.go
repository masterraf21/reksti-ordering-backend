package usecases

import "github.com/masterraf21/reksti-ordering-backend/models"

type orderUsecase struct {
	orderRepo        models.OrderRepository
	orderDetailsRepo models.OrderDetailsRepository
}

// NewOrderUsecase will create new order usecase
func NewOrderUsecase(
	ouc models.OrderRepository,
	odc models.OrderDetailsRepository,
) models.OrderUsecase {
	return &orderUsecase{
		orderRepo:        ouc,
		orderDetailsRepo: odc,
	}
}
