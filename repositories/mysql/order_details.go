package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type orderDetailsRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewOrderDetailsRepo create new order repo
func NewOrderDetailsRepo(reader, writer *sql.DB) models.OrderDetailsRepository {
	return &orderDetailsRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *orderDetailsRepo) UpdateTotalPrice(ctx context.Context, orderDetailsID uint32) error {
	orderDetail, err := r.GetByID(orderDetailsID)
	if err != nil {
		return err
	}

	var menu models.Menu
	menuQuery := sq.Select("*").
		From("menu").
		Where(
			sq.Eq{
				"menu_id": orderDetail.MenuID,
			},
		).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)
	rows, err := menuQuery.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&menu.MenuID,
			&menu.Name,
			&menu.Price,
			&menu.MenuTypeID,
			&menu.Ingredients,
			&menu.MenuStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
	}

	newPrice := float32(orderDetail.Quantity) * menu.Price
	err = r.UpdateArbitrary(
		ctx,
		orderDetailsID,
		"total_price",
		newPrice,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) GetOrderDetailsByOrderID(
	orderID uint32,
) (res []models.OrderDetails, err error) {
	table := "order_details"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"order_id": orderID,
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
		var r models.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.MenuID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderDetailsRepo) GetAll() (res []models.OrderDetails, err error) {
	table := "order_details"

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
		var r models.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.MenuID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderDetailsRepo) GetByID(orderDetailsID uint32) (res *models.OrderDetails, err error) {
	table := "order_details"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"order_details_id": orderDetailsID,
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
		var r models.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.MenuID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *orderDetailsRepo) UpdateArbitrary(
	ctx context.Context,
	orderDetailsID uint32,
	columnName string,
	value interface{},
) error {
	table := "order_details"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_details_id": orderDetailsID,
		}).
		Set(columnName, value).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) UpdateByID(
	ctx context.Context,
	orderDetailsID uint32,
	order *models.OrderDetails,
) error {
	table := "order_details"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_details_id": orderDetailsID,
		}).
		SetMap(map[string]interface{}{
			"order_id":    order.OrderID,
			"menu_id":     order.MenuID,
			"quantity":    order.Quantity,
			"total_price": order.TotalPrice,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) DeleteByID(ctx context.Context, orderDetailsID uint32) error {
	table := "order_details"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"order_details_id": orderDetailsID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderDetailsRepo) Store(
	ctx context.Context,
	ord *models.OrderDetails,
) (orderDetailsID uint32, err error) {
	table := "order_details"
	query := sq.Insert(table).
		Columns(
			"order_id",
			"menu_id",
			"quantity",
			"total_price",
		).
		Values(
			ord.OrderID,
			ord.MenuID,
			ord.Quantity,
			ord.TotalPrice,
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
	orderDetailsID = uint32(id)

	return
}

func (r *orderDetailsRepo) BulkInsert(
	ctx context.Context,
	orderDetails []models.OrderDetails,
) error {
	table := "order_details"
	tx, err := r.Writer.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var insertQuery sq.InsertBuilder
	for _, ord := range orderDetails {
		insertQuery = sq.Insert(table).
			Columns(
				"order_id",
				"menu_id",
				"quantity",
				"total_price",
			).
			Values(
				ord.OrderID,
				ord.MenuID,
				ord.Quantity,
				ord.TotalPrice,
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
