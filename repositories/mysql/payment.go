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
			&r.PaymentType,
			&r.PaymentDate,
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
			&r.PaymentType,
			&r.PaymentDate,
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

func (t *paymentRepo) Store(ctx context.Context, payment *models.Payment) error {
	table := "payment"
	now := time.Now()
	nowInsert := now.Format(time.RFC3339)

	query := sq.Insert(table).
		Columns(
			"order_id",
			"amount",
			"payment_type",
			"payment_date",
		).
		Values(
			payment.OrderID,
			payment.Amount,
			payment.PaymentType,
			nowInsert,
		)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return err
}
