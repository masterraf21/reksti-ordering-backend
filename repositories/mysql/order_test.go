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
	Reader     *sql.DB
	Writer     *sql.DB
	Repo       models.OrderRepository
	DetailRepo models.OrderDetailsRepository
	CustomerID uint32
	MenuID     uint32
	MenuPrice  float32
}

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(orderRepoTestSuite))
}

func (s *orderRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewOrderRepo(reader, writer)
	s.DetailRepo = NewOrderDetailsRepo(reader, writer)
	customerRepo := NewCustomerRepo(reader, writer)
	menuRepo := NewMenuRepo(reader, writer)
	menuTypeRepo := NewMenuTypeRepo(reader, writer)

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

	menuType := models.MenuType{
		TypeName:    "Hihi",
		Description: "Yuhu",
	}
	mtID, err := menuTypeRepo.Store(context.TODO(), &menuType)
	if err != nil {
		panic(err)
	}

	menu := models.Menu{
		MenuTypeID:  mtID,
		Name:        "Pokemon",
		Price:       10000,
		Ingredients: "Pikachoe",
		MenuStatus:  true,
	}
	mID, err := menuRepo.Store(context.TODO(), &menu)
	if err != nil {
		panic(err)
	}
	s.MenuID = mID
	s.MenuPrice = menu.Price
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
		order := models.Order{
			CustomerID:  s.CustomerID,
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

func (s *orderRepoTestSuite) TestGet() {
	s.Run("Get Order By Id", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
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

	s.Run("Get Order by Status and Customer ID", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 1,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		results, err := s.Repo.GetByStatusAndCustID([]int32{order.OrderStatus}, s.CustomerID)
		result := results[0]
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(1, len(results))
		s.Equal(order.CustomerID, result.CustomerID)
		s.Equal(id, result.OrderID)
		s.Equal(order.TotalPrice, result.TotalPrice)
		s.Equal(order.OrderStatus, result.OrderStatus)
	})

	s.Run("Get Order by Customer ID", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		results, err := s.Repo.GetByCustID(s.CustomerID)
		result := results[0]
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(3, len(results))
		s.Equal(order.CustomerID, result.CustomerID)
		s.Equal(order.TotalPrice, result.TotalPrice)
		s.Equal(order.OrderStatus, result.OrderStatus)
	})
}

func (s *orderRepoTestSuite) TestGet2() {
	s.Run("Get All Orders", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(context.TODO(), &order)
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

func (s *orderRepoTestSuite) TestUpdate() {
	s.Run("Update Order Status", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		err = s.Repo.UpdateOrderStatus(context.TODO(), id, 2)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().EqualValues(2, result.OrderStatus)
	})

	s.Run("Update Order Total Price", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		countDetail := 3
		detail := models.OrderDetails{
			OrderID:    id,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}
		for i := 0; i < countDetail; i++ {
			id, err := s.DetailRepo.Store(context.TODO(), &detail)
			if err != nil {
				panic(err)
			}
			err = s.DetailRepo.UpdateTotalPrice(context.TODO(), id)
			if err != nil {
				panic(err)
			}
		}
		expectedPrice := s.MenuPrice * float32(detail.Quantity) * float32(countDetail)

		err = s.Repo.UpdateTotalPrice(context.TODO(), id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(expectedPrice, result.TotalPrice)
	})
}

func (s *orderRepoTestSuite) TestDelete() {
	s.Run("Delete Order By ID", func() {
		order := models.Order{
			CustomerID:  s.CustomerID,
			TotalPrice:  0,
			OrderStatus: 0,
		}
		id, err := s.Repo.Store(context.TODO(), &order)
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

		s.Assert().Nil(result)
	})
}
