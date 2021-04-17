package mysql

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/masterraf21/reksti-ordering-backend/models"
	logger "github.com/sirupsen/logrus"
)

type paymentRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewPaymentRepo will initiate repo
func NewPaymentRepo(reader, writer *sql.DB) models.PaymentRepository {
	return &paymentRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (t *paymentRepo) UpdateStatus(ctx context.Context, paymentID uint32, status int32) error {
	table := "payment"

	query := sq.Update(table).
		Set("payment_status", status).
		RunWith(t.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *paymentRepo) GetAll() (res []models.Payment, err error) {
	table := "payment"

	query := sq.Select("*").
		From(table).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Payment
		err = rows.Scan(
			&r.PaymentID,
			&r.OrderID,
			&r.Amount,
			&r.PaymentTypeID,
			&r.PaymentDate,
			&r.PaymentStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (t *paymentRepo) GetByID(paymentID uint32) (res *models.Payment, err error) {
	table := "payment"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"payment_id": paymentID,
		}).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Payment
		err = rows.Scan(
			&r.PaymentID,
			&r.OrderID,
			&r.Amount,
			&r.PaymentTypeID,
			&r.PaymentDate,
			&r.PaymentStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (t *paymentRepo) DeleteByID(ctx context.Context, paymentID uint32) error {
	table := "payment"

	query := sq.Delete("").
		From(table).
		Where(sq.Eq{
			"payment_id": paymentID,
		}).
		RunWith(t.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *paymentRepo) Store(
	ctx context.Context,
	payment *models.Payment,
) (paymentID uint32, err error) {
	table := "payment"
	now := time.Now()
	nowInsert := now.Format(time.RFC3339)

	query := sq.Insert(table).
		Columns(
			"order_id",
			"amount",
			"payment_type_id",
			"payment_date",
		).
		Values(
			payment.OrderID,
			payment.Amount,
			payment.PaymentTypeID,
			nowInsert,
		).RunWith(t.Writer).
		PlaceholderFormat(sq.Question)
	sqlInsert, argsInsert, err := query.ToSql()
	res, err := t.Writer.ExecContext(
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
	paymentID = uint32(id)

	return
}
