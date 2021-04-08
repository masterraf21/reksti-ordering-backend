package mysql

import (
	"context"
	"database/sql"

	logger "github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type orderRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewOrderRepo create new order repo
func NewOrderRepo(reader, writer *sql.DB) models.OrderRepository {
	return &orderRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *orderRepo) GetAll() (res []models.Order, err error) {
	table := "orders"

	query := sq.Select("*").
		From(table).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Order
		err = rows.Scan(
			&r.OrderID,
			&r.CustomerID,
			&r.OrderDate,
			&r.TotalAmount,
			&r.OrderStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderRepo) GetByID(OrderID uint32) (res *models.Order, err error) {
	table := "orders"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"order_id": OrderID,
			},
		).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Order
		err = rows.Scan(
			&r.OrderID,
			&r.CustomerID,
			&r.OrderDate,
			&r.TotalAmount,
			&r.OrderStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *orderRepo) Update(
	ctx context.Context,
	order *models.Order) error

func (r *orderRepo) DeleteByID(
	ctx context.Context,
	OrderID uint32,
) error {
	table := "orders"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"order_id": OrderID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepo) Store(
	ctx context.Context,
	ord *models.Order) error
