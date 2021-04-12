package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type customerRepoTestSuite struct {
	suite.Suite
	Reader *sql.DB
	Writer *sql.DB
	Repo   models.CustomerRepository
}

func TestCustomerRepository(t *testing.T) {
	suite.Run(t, new(customerRepoTestSuite))
}

func (s *customerRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewCustomerRepo(reader, writer)
}

func (s *customerRepoTestSuite) TearDownSuite() {
	querys := []string{
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

func (s *customerRepoTestSuite) TearDownTest() {
	querys := []string{
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

func (s *customerRepoTestSuite) TestStore() {
	s.Run("Store Customer", func() {
		customer := models.Customer{
			FullName:    "name test",
			Email:       "email test",
			PhoneNumber: "081xxxxxx",
			Username:    "ohayo poko",
			Password:    "return_to_monke",
		}
		id, err := s.Repo.Store(context.TODO(), &customer)
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
		s.Equal(id, result.CustomerID)
		s.Equal(customer.FullName, result.FullName)
		s.Equal(customer.Email, result.Email)
		s.Equal(customer.PhoneNumber, result.PhoneNumber)
		s.Equal(customer.Username, result.Username)
		s.Equal(customer.Password, result.Password)
	})
}

func (s *customerRepoTestSuite) TestGet1() {
	s.Run("Get Customer By ID", func() {
		customer := models.Customer{
			FullName:    "name test",
			Email:       "email test",
			PhoneNumber: "081xxxxxx",
			Username:    "ohayo poko",
			Password:    "return_to_monke",
		}
		id, err := s.Repo.Store(context.TODO(), &customer)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(id, result.CustomerID)
		s.Equal(customer.FullName, result.FullName)
		s.Equal(customer.Email, result.Email)
		s.Equal(customer.PhoneNumber, result.PhoneNumber)
		s.Equal(customer.Username, result.Username)
		s.Equal(customer.Password, result.Password)
	})
}

func (s *customerRepoTestSuite) TestGet2() {
	s.Run("Get All Customer", func() {
		customer := models.Customer{
			FullName:    "name test",
			Email:       "email test",
			PhoneNumber: "081xxxxxx",
			Username:    "ohayo poko",
			Password:    "return_to_monke",
		}
		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(context.TODO(), &customer)
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

func (s *customerRepoTestSuite) TestDelete() {
	s.Run("Delete Customer", func() {
		customer := models.Customer{
			FullName:    "name test",
			Email:       "email test",
			PhoneNumber: "081xxxxxx",
			Username:    "ohayo poko",
			Password:    "return_to_monke",
		}
		id, err := s.Repo.Store(context.TODO(), &customer)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

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
