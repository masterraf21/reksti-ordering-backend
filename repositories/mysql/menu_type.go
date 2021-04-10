package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type menuTypeRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewOrderDetailsRepo create new order repo
func NewMenuTypeRepo(reader, writer *sql.DB) models.MenuTypeRepository {
	return &menuTypeRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *menuTypeRepo) GetAll() (res []models.MenuType, err error) {
	table := "menu_type"

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
		var r models.MenuType
		err = rows.Scan(
			&r.MenuTypeID,
			&r.TypeName,
			&r.Description,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *menuTypeRepo) GetByID(menuTypeID uint32) (res *models.MenuType, err error) {
	table := "menu_type"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"menu_type_id": menuTypeID,
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
		var r models.MenuType
		err = rows.Scan(
			&r.MenuTypeID,
			&r.TypeName,
			&r.Description,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *menuTypeRepo) UpdateArbitrary(
	ctx context.Context,
	menuTypeID uint32,
	columnName string,
	value interface{},
) error {
	table := "menu_type"
	query := sq.Update(table).
		Where(sq.Eq{
			"menu_type_id": menuTypeID,
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

func (r *menuTypeRepo) UpdateByID(
	ctx context.Context,
	menuTypeID uint32,
	order *models.MenuType,
) error {
	table := "menu_type"
	query := sq.Update(table).
		Where(sq.Eq{
			"menu_type_id": menuTypeID,
		}).
		SetMap(map[string]interface{}{
			"type_name":    order.TypeName,
			"description":     order.Description,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *menuTypeRepo) DeleteByID(ctx context.Context, menuTypeID uint32) error {
	table := "menu_type"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"menu_type_id": menuTypeID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *menuTypeRepo) Store(
	ctx context.Context,
	ord *models.MenuType,
) (menuTypeID uint32, err error) {
	table := "menu_type"
	query := sq.Insert(table).
		Columns(
			"type_name",
			"description",
		).
		Values(
			ord.TypeName,
			ord.Description,
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
	menuTypeID = uint32(id)

	return
}

func (r *menuTypeRepo) BulkInsert(
	ctx context.Context,
	MenuType []models.MenuType,
) error {
	table := "menu_type"
	tx, err := r.Writer.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var insertQuery sq.InsertBuilder
	for _, ord := range MenuType {
		insertQuery = sq.Insert(table).
			Columns(
				"type_name",
				"description",
			).
			Values(
				ord.TypeName,
				ord.Description,
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
