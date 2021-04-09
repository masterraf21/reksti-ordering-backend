package usecases

import (
	"context"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

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

func (u *orderUsecase) GetOrderWithDetails(orderID uint32) (res models.OrderWithDetails, err error) {
	order, err := u.orderRepo.GetByID(orderID)
	if err != nil {
		return
	}

	orderDetails, err := u.orderRepo.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return
	}

	res.Details = orderDetails
	res.OrderInfo = *order

	return
}

func (u *orderUsecase) GetAllOrders() (res []models.Order, err error) {
	res, err = u.orderRepo.GetAll()
	return
}

func (u *orderUsecase) GetOrderByID(orderID uint32) (res *models.Order, err error) {
	res, err = u.orderRepo.GetByID(orderID)
	return
}

func (u *orderUsecase) GetOrdersByCustID(customerID uint32) (res []models.Order, err error) {
	res, err = u.orderRepo.GetByCustID(customerID)
	return
}

func (u *orderUsecase) GetOrdersHistoryByCustID(customerID uint32) (res []models.Order, err error) {
	res, err = u.orderRepo.GetByStatusAndCustID(
		int32(2),
		customerID,
	)
	return
}

func (u *orderUsecase) GetOngoingOrdersyByCustID(customerID uint32) (res []models.Order, err error) {
	res, err = u.orderRepo.GetByStatusAndCustID(
		int32(0),
		customerID,
	)
	return
}

func (u *orderUsecase) CreateOrder(ctx context.Context, order *models.Order) (id uint32, err error) {
	id, err = u.CreateOrder(
		ctx,
		order,
	)
	return
}

func (u *orderUsecase) BulkCreateOrders(ctx context.Context, orders []models.Order) error {
	err := u.orderRepo.BulkInsert(
		ctx,
		orders,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) UpdateOrder(ctx context.Context, orderID uint32, order *models.Order) error {
	err := u.orderRepo.UpdateByID(
		ctx,
		orderID,
		order,
	)
	if err != nil {
		return err
	}

	return nil
}
