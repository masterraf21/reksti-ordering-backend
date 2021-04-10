package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type menuRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewMenuRepo create new order repo
func NewMenuRepo(reader, writer *sql.DB) models.MenuRepository {
	return &menuRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *menuRepo) GetAll() (res []models.Menu, err error) {
	table := "menu"

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
		var r models.Menu
		err = rows.Scan(
			&r.MenuID,
			&r.Name,
			&r.Price,
			&r.MenuTypeID,
			&r.Ingredients,
			&r.MenuStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *menuRepo) GetByID(menuID uint32) (res *models.Menu, err error) {
	table := "menu"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"menu_id": menuID,
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
		var r models.Menu
		err = rows.Scan(
			&r.MenuID,
			&r.Name,
			&r.Price,
			&r.MenuTypeID,
			&r.Ingredients,
			&r.MenuStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *menuRepo) UpdateArbitrary(
	ctx context.Context,
	menuID uint32,
	columnName string,
	value interface{},
) error {
	table := "menu"
	query := sq.Update(table).
		Where(sq.Eq{
			"menu_id": menuID,
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

func (r *menuRepo) UpdateByID(
	ctx context.Context,
	menuID uint32,
	order *models.Menu,
) error {
	table := "menu"
	query := sq.Update(table).
		Where(sq.Eq{
			"menu_id": menuID,
		}).
		SetMap(map[string]interface{}{
			"menu_name":    order.Name,
			"price":     order.Price,
			"menu_type_id":     order.MenuTypeID,
			"ingredients":     order.Ingredients,
			"menu_status":     order.MenuStatus,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *menuRepo) DeleteByID(ctx context.Context, menuID uint32) error {
	table := "menu"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"menu_id": menuID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *menuRepo) Store(
	ctx context.Context,
	ord *models.Menu,
) (menuID uint32, err error) {
	table := "menu"
	query := sq.Insert(table).
		Columns(
			"menu_name",
			"price",
			"menu_type_id",
			"ingredients",
			"menu_status",
		).
		Values(
			ord.Name,
			ord.Price,
			ord.MenuTypeID,
			ord.Ingredients,
			ord.MenuStatus,
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
	menuID = uint32(id)

	return
}

func (r *menuRepo) BulkInsert(
	ctx context.Context,
	Menu []models.Menu,
) error {
	table := "menu"
	tx, err := r.Writer.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var insertQuery sq.InsertBuilder
	for _, ord := range Menu {
		insertQuery = sq.Insert(table).
			Columns(
				"menu_name",
				"price",
				"menu_type_id",
				"ingredients",
				"menu_status",
			).
			Values(
				ord.Name,
				ord.Price,
				ord.MenuTypeID,
				ord.Ingredients,
				ord.MenuStatus,
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
