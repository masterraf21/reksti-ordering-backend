package mysql

import (
	"database/sql"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type orderDetailsRepoTestSuite struct {
	suite.Suite
	Reader *sql.DB
	Writer *sql.DB
	Repo   models.OrderDetailsRepository
}
