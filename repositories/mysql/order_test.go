package mysql

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/configs"
	"github.com/masterraf21/reksti-ordering-backend/models"
	"github.com/masterraf21/reksti-ordering-backend/utils/mysql"
)

func configureMySQL() (*sql.DB, *sql.DB) {
	readerConfig := mysql.Option{
		Host:     configs.MySQL.ReaderHost,
		Port:     configs.MySQL.ReaderPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.ReaderUser,
		Password: configs.MySQL.ReaderPassword,
	}

	writerConfig := mysql.Option{
		Host:     configs.MySQL.WriterHost,
		Port:     configs.MySQL.WriterPort,
		Database: configs.MySQL.Database,
		User:     configs.MySQL.WriterUser,
		Password: configs.MySQL.WriterPassword,
	}

	reader, writer, err := mysql.SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}

	log.Println("MySQL connection is successfully established!")

	return reader, writer
}

type orderRepoTestSuite struct {
	suite.Suite
	Reader       *sql.DB
	Writer       *sql.DB
	Repo         models.OrderRepository
	CustomerRepo models.CustomerRepository
}

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(orderRepoTestSuite))
}

func (s *orderRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewOrderRepo(reader, writer)
	s.CustomerRepo = NewCustomerRepo(reader, writer)
}

func (s *orderRepoTestSuite) TearDownSuite() {
	querys := []string{
		"DELETE FROM orders;",
		"DELETE FROM customer;",
	}
	var err error
	for _, query := range querys {
		_, err = s.Writer.ExecContext(
			context.TODO(),
			query,
		)
		if err != nil {
			panic(err)
		}
	}
}

func (s *orderRepoTestSuite) TearDownTest() {
	querys := []string{
		"DELETE FROM orders;",
		"DELETE FROM customer;",
	}
	var err error
	for _, query := range querys {
		_, err = s.Writer.ExecContext(
			context.TODO(),
			query,
		)
		if err != nil {
			panic(err)
		}
	}
}

func (s *orderRepoTestSuite) TestInserts() {
	s.Run("Store Order", func() {
		cust := models.Customer{
			FullName:    "name test",
			Email:       "email test",
			PhoneNumber: "081xxxxxx",
			Username:    "ohayo poko",
			Password:    "return_to_monke",
		}
		custID, err := s.CustomerRepo.Store(context.TODO(), &cust)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		order := models.Order{
			CustomerID:  custID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(order.CustomerID, result.CustomerID)
		s.Equal(id, result.OrderID)
		s.Equal(order.TotalPrice, result.TotalPrice)
		s.Equal(order.OrderStatus, result.OrderStatus)
	})
}
