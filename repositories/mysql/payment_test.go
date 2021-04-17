package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/masterraf21/reksti-ordering-backend/models"
	"github.com/stretchr/testify/suite"
)

type paymentRepoTestSuite struct {
	suite.Suite
	Reader        *sql.DB
	Writer        *sql.DB
	Repo          models.PaymentRepository
	OrderID       uint32
	CustomerID    uint32
	PaymentTypeID uint32
}

func TestPaymentRepository(t *testing.T) {
	suite.Run(t, new(paymentRepoTestSuite))
}

func (s *paymentRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewPaymentRepo(reader, writer)
	orderRepo := NewOrderRepo(reader, writer)
	customerRepo := NewCustomerRepo(reader, writer)
	paymentTypeRepo := NewPaymentTypeRepo(reader, writer)

	cust := models.Customer{
		FullName:    "name test",
		Email:       "email test",
		PhoneNumber: "081xxxxxx",
		Username:    "ohayo poko",
		Password:    "return_to_monke",
	}
	custID, err := customerRepo.Store(context.TODO(), &cust)
	if err != nil {
		panic(err)
	}
	s.CustomerID = custID

	order := models.Order{
		CustomerID:  custID,
		TotalPrice:  0,
		OrderStatus: 0,
	}
	orderID, err := orderRepo.Store(context.TODO(), &order)
	if err != nil {
		panic(err)
	}
	s.OrderID = orderID

	paymentType := models.PaymentType{
		Method:  "GoPay",
		Company: "GoJek",
	}
	ptID, err := paymentTypeRepo.Store(context.TODO(), &paymentType)
	if err != nil {
		panic(err)
	}
	s.PaymentTypeID = ptID
}

func (s *paymentRepoTestSuite) TearDownSuite() {
	querys := []string{
		"DELETE FROM payment;",
		"DELETE FROM orders;",
		"DELETE FROM customer;",
		"DELETE FROM payment_type;",
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

func (s *paymentRepoTestSuite) TearDownTest() {
	querys := []string{
		"DELETE FROM payment;",
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

func (s *paymentRepoTestSuite) TestStore() {
	s.Run("Store Payment", func() {
		payment := models.Payment{
			OrderID:       s.OrderID,
			Amount:        10000,
			PaymentTypeID: s.PaymentTypeID,
		}
		id, err := s.Repo.Store(context.TODO(), &payment)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		result, err := s.Repo.GetAll()
		a := result[0]
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(1, len(result))
		s.Equal(payment.OrderID, a.OrderID)
		s.Equal(payment.Amount, a.Amount)
		s.Equal(payment.PaymentTypeID, a.PaymentTypeID)
	})
}

func (s *paymentRepoTestSuite) TestGet() {
	s.Run("Get All Payment", func() {
		payment := models.Payment{
			OrderID:       s.OrderID,
			Amount:        10000,
			PaymentTypeID: s.PaymentTypeID,
		}

		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(context.TODO(), &payment)
			if err != nil {
				panic(err)
			}
			s.Require().NoError(err)
			s.Assert().NotEmpty(id)
		}

		result, err := s.Repo.GetAll()
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(5, len(result))
	})
}

func (s *paymentRepoTestSuite) TestGet2() {
	s.Run("Get Payment By ID", func() {
		payment := models.Payment{
			OrderID:       s.OrderID,
			Amount:        10000,
			PaymentTypeID: s.PaymentTypeID,
		}

		id, err := s.Repo.Store(context.TODO(), &payment)
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
		s.Equal(payment.Amount, result.Amount)
		s.Equal(payment.OrderID, result.OrderID)
		s.Equal(id, result.PaymentID)
		s.Equal(payment.PaymentTypeID, result.PaymentTypeID)
	})
}

func (s *paymentRepoTestSuite) TestDelete() {
	s.Run("Delete Payment By ID", func() {
		payment := models.Payment{
			OrderID:       s.OrderID,
			Amount:        10000,
			PaymentTypeID: s.PaymentTypeID,
		}

		id, err := s.Repo.Store(context.TODO(), &payment)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		err = s.Repo.DeleteByID(context.TODO(), id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Nil(result)
	})
}
