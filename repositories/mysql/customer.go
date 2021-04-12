package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/masterraf21/reksti-ordering-backend/models"
	logger "github.com/sirupsen/logrus"
)

type customerRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewCustomerRepo will create customer rrepo
func NewCustomerRepo(reader, writer *sql.DB) models.CustomerRepository {
	return &customerRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *customerRepo) GetAll() (res []models.Customer, err error) {
	table := "customer"

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
		var r models.Customer
		err = rows.Scan(
			&r.CustomerID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.AccountStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *customerRepo) GetByID(CustomerID uint32) (res *models.Customer, err error) {
	table := "customer"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"customer_id": CustomerID,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Customer
		err = rows.Scan(
			&r.CustomerID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.AccountStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *customerRepo) DeleteByID(ctx context.Context, CustomerID uint32) error {
	table := "customer"

	query := sq.Delete("").
		From(table).
		Where(sq.Eq{
			"customer_id": CustomerID,
		}).
		RunWith(r.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepo) Store(ctx context.Context, cust *models.Customer) error {
	table := "customer"

	query := sq.Insert(table).
		Columns(
			"customer_full_name",
			"customer_email",
			"customer_phone_number",
			"customer_username",
			"customer_password",
			"account_status",
		).
		Values(
			cust.FullName,
			cust.Email,
			cust.PhoneNumber,
			cust.Username,
			cust.Password,
			cust.AccountStatus,
		).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
