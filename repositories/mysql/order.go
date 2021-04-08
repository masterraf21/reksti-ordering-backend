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

func (r *orderRepo) UpdateTotalPrice(
	ctx context.Context,
	orderID uint32,
) error {
}

func (r *orderRepo) UpdateOrderStatus(
	ctx context.Context,
	orderID uint32,
	status int32,
) error {
	table := "orders"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_id": orderID,
		}).
		Set("order_status", status).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepo) GetByStatusAndCustID(
	status int32,
	custID uint32,
) (res []models.Order, err error) {
	table := "orders"

	andStatement := sq.And{
		sq.Eq{
			"customer_id": custID,
		},
		sq.Eq{
			"order_status": status,
		},
	}
	query := sq.Select("*").
		From(table).
		Where(andStatement).
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
			&r.TotalPrice,
			&r.OrderStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderRepo) GetByCustID(custID uint32) (res []models.Order, err error) {
	table := "orders"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"customer_id": custID,
		}).
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
			&r.TotalPrice,
			&r.OrderStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
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
			&r.TotalPrice,
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
			&r.TotalPrice,
			&r.OrderStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *orderRepo) UpdateByID(
	ctx context.Context,
	orderID uint32,
	order *models.Order,
) error {
	table := "orders"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_id": orderID,
		}).
		SetMap(map[string]interface{}{
			"customer_id":  order.CustomerID,
			"order_date":   order.OrderDate,
			"total_price":  order.TotalPrice,
			"order_status": order.OrderStatus,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

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
	ord *models.Order,
) (orderID uint32, err error) {
	table := "orders"
	query := sq.Insert(table).
		Columns(
			"customer_id",
			"order_date",
			"total_price",
			"order_status",
		).
		Values(
			ord.CustomerID,
			ord.OrderDate,
			ord.TotalPrice,
			ord.OrderStatus,
		).
		PlaceholderFormat(sq.Question)
	sqlInsert, argsInsert, err := query.ToSql()
	res, err := r.Writer.ExecContext(
		ctx,
		sqlInsert,
		argsInsert...,
	)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	orderID = uint32(id)

	return
}

func (r *orderRepo) BulkInsert(
	ctx context.Context,
	orders []models.Order,
) error {
	table := "orders"
	tx, err := r.Writer.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var insertQuery sq.InsertBuilder
	for _, ord := range orders {
		insertQuery = sq.Insert(table).
			Columns(
				"customer_id",
				"order_date",
				"total_price",
				"order_status",
			).
			Values(
				ord.CustomerID,
				ord.OrderDate,
				ord.TotalPrice,
				ord.OrderStatus,
			).
			PlaceholderFormat(sq.Question)
		sqlInsert, argsInsert, err := insertQuery.ToSql()
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(
			ctx,
			sqlInsert,
			argsInsert...,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
