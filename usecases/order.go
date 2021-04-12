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

func (u *orderUsecase) GetOrderWithDetails(
	orderID uint32,
) (res models.OrderWithDetails, err error) {
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

func (u *orderUsecase) GetOrderDetailByID(orderDetailID uint32) (res *models.OrderDetails, err error) {
	res, err = u.orderDetailsRepo.GetByID(orderDetailID)
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

func (u *orderUsecase) GetOngoingOrdersyByCustID(
	customerID uint32,
) (res []models.Order, err error) {
	res, err = u.orderRepo.GetByStatusAndCustID(
		int32(0),
		customerID,
	)
	return
}

func (u *orderUsecase) CreateOrder(
	ctx context.Context,
	order *models.Order,
) (id uint32, err error) {
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

func (u *orderUsecase) CreateOrderDetail(
	ctx context.Context,
	orderD *models.OrderDetails,
) error {
	err := u.orderDetailsRepo.Store(
		ctx,
		orderD,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) UpdateOrderPrice(ctx context.Context, orderID uint32) error {
	err := u.orderRepo.UpdateTotalPrice(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) UpdateOrderDetailPrice(ctx context.Context, orderDetailID uint32) error {
	err := u.orderDetailsRepo.UpdateTotalPrice(ctx, orderDetailID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) DeleteOrder(ctx context.Context, orderID uint32) error {
	err := u.orderRepo.DeleteByID(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) DeleteOrderDetail(ctx context.Context, orderDetailID uint32) error {
	err := u.orderDetailsRepo.DeleteByID(ctx, orderDetailID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) UpdateOrderStatus(ctx context.Context, orderID uint32, status int32) error {
	err := u.orderRepo.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		return err
	}

	return nil
}
