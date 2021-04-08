package mysql

import (
	"database/sql"

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
