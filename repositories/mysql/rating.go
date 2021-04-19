package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type ratingRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewRatingRepo create new order repo
func NewRatingRepo(reader, writer *sql.DB) models.RatingRepository {
	return &ratingRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *ratingRepo) GetAll() (res []models.Rating, err error) {
	table := "rating"

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
		var r models.Rating
		err = rows.Scan(
			&r.RatingID,
			&r.MenuID,
			&r.Score,
			&r.Remarks,
			&r.DateRecorded,
			&r.CustomerID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *ratingRepo) GetByID(ratingID uint32) (res *models.Rating, err error) {
	table := "rating"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"rating_id": ratingID,
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
		var r models.Rating
		err = rows.Scan(
			&r.RatingID,
			&r.MenuID,
			&r.Score,
			&r.Remarks,
			&r.DateRecorded,
			&r.CustomerID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *ratingRepo) GetByMenu(menuID uint32) (res []models.Rating, err error) {
	table := "rating"

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
		var r models.Rating
		err = rows.Scan(
			&r.RatingID,
			&r.MenuID,
			&r.Score,
			&r.Remarks,
			&r.DateRecorded,
			&r.CustomerID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *ratingRepo) GetMenuScore(menuID uint32) (res float32, err error) {
	table := "rating"

	query := sq.Select("score").
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

	var sum float32
	count := 0
	for rows.Next() {
		var score float32
		err = rows.Scan(
			&score,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		sum += score
		count++
	}
	res = sum / float32(count)

	return
}

func (r *ratingRepo) UpdateArbitrary(
	ratingID uint32,
	columnName string,
	value interface{},
) error {
	table := "rating"
	query := sq.Update(table).
		Where(sq.Eq{
			"rating_id": ratingID,
		}).
		Set(columnName, value).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *ratingRepo) UpdateByID(
	ratingID uint32,
	order *models.Rating,
) error {
	table := "rating"
	query := sq.Update(table).
		Where(sq.Eq{
			"rating_id": ratingID,
		}).
		SetMap(map[string]interface{}{
			"menu_id":       order.MenuID,
			"score":         order.Score,
			"remarks":       order.Remarks,
			"date_recorded": order.DateRecorded,
			"customer_id":   order.CustomerID,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *ratingRepo) DeleteByID(ratingID uint32) error {
	table := "rating"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"rating_id": ratingID,
			},
		).
		RunWith(r.Writer)
	_, err := query.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *ratingRepo) Store(
	ord *models.Rating,
) (ratingID uint32, err error) {
	table := "rating"
	query := sq.Insert(table).
		Columns(
			"menu_id",
			"score",
			"remarks",
			"date_recorded",
			"customer_id",
		).
		Values(
			ord.MenuID,
			ord.Score,
			ord.Remarks,
			sq.Expr("NOW()"),
			ord.CustomerID,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	res, err := r.Writer.Exec(
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
	ratingID = uint32(id)

	return
}

func (r *ratingRepo) BulkInsert(
	Rating []models.Rating,
) error {
	table := "rating"
	tx, err := r.Writer.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	var insertQuery sq.InsertBuilder
	for _, ord := range Rating {
		insertQuery = sq.Insert(table).
			Columns(
				"menu_id",
				"score",
				"remarks",
				"date_recorded",
				"customer_id",
			).
			Values(
				ord.MenuID,
				ord.Score,
				ord.Remarks,
				sq.Expr("NOW()"),
				ord.CustomerID,
			).
			PlaceholderFormat(sq.Question)

		sqlInsert, argsInsert, err := insertQuery.ToSql()
		if err != nil {
			return err
		}
		_, err = tx.Exec(
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
