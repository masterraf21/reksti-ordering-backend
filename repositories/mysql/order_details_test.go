package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type orderDetailsRepoTestSuite struct {
	suite.Suite
	Reader    *sql.DB
	Writer    *sql.DB
	Repo      models.OrderDetailsRepository
	OrderID   uint32
	MenuID    uint32
	MenuPrice float32
}

func TestOrderDetailRepository(t *testing.T) {
	suite.Run(t, new(orderDetailsRepoTestSuite))
}

func (s *orderDetailsRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewOrderDetailsRepo(reader, writer)

	orderRepo := NewOrderRepo(reader, writer)
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
	order := models.Order{
		CustomerID:  custID,
		TotalPrice:  0,
		OrderStatus: 0,
	}
	orderID, err := orderRepo.Store(context.TODO(), &order)
	if err != nil {
		panic(err)
	}

	menuType := models.MenuType{
		TypeName:    "Yehe",
		Description: "Yuhu",
	}
	menuTypeID, err := menuTypeRepo.Store(context.TODO(), &menuType)
	if err != nil {
		panic(err)
	}

	menu := models.Menu{
		MenuTypeID:  menuTypeID,
		Name:        "Siomay",
		Price:       10000,
		Ingredients: "Tahu Bakso Bulls",
		MenuStatus:  true,
	}
	menuID, err := menuRepo.Store(context.TODO(), &menu)
	if err != nil {
		panic(err)
	}

	s.OrderID = orderID
	s.MenuID = menuID
	s.MenuPrice = menu.Price
}

func (s *orderDetailsRepoTestSuite) TearDownSuite() {
	querys := []string{
		"DELETE FROM order_details;",
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

func (s *orderDetailsRepoTestSuite) TearDownTest() {
	querys := []string{
		"DELETE FROM order_details;",
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

func (s *orderDetailsRepoTestSuite) TestStore() {
	s.Run("Store Order Details", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}

		id, err := s.Repo.Store(context.TODO(), &ordDetail)
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
		s.Equal(id, result.OrderDetailsID)
		s.Equal(ordDetail.OrderID, result.OrderID)
		s.Equal(ordDetail.MenuID, result.MenuID)
		s.Equal(ordDetail.Quantity, result.Quantity)
		s.Equal(ordDetail.TotalPrice, result.TotalPrice)
	})
}

func (s *orderDetailsRepoTestSuite) TestGet() {
	s.Run("Get Order Details By ID", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}

		id, err := s.Repo.Store(context.TODO(), &ordDetail)
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
		s.Equal(id, result.OrderDetailsID)
		s.Equal(ordDetail.OrderID, result.OrderID)
		s.Equal(ordDetail.MenuID, result.MenuID)
		s.Equal(ordDetail.Quantity, result.Quantity)
		s.Equal(ordDetail.TotalPrice, result.TotalPrice)
	})

	s.Run("Get Order Details By Order ID", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}

		id, err := s.Repo.Store(context.TODO(), &ordDetail)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		result, err := s.Repo.GetOrderDetailsByOrderID(s.OrderID)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(2, len(result))
	})
}

func (s *orderDetailsRepoTestSuite) TestGet2() {
	s.Run("Get All Order Details", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}
		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(context.TODO(), &ordDetail)
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

func (s *orderDetailsRepoTestSuite) TestUpdate() {
	s.Run("Update Order Detail Total Price", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}

		id, err := s.Repo.Store(context.TODO(), &ordDetail)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		expectedPrice := s.MenuPrice * float32(ordDetail.Quantity)

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

func (s *orderDetailsRepoTestSuite) TestDelete() {
	s.Run("Delete Order Detail by ID", func() {
		ordDetail := models.OrderDetails{
			OrderID:    s.OrderID,
			MenuID:     s.MenuID,
			Quantity:   10,
			TotalPrice: 0,
		}

		id, err := s.Repo.Store(context.TODO(), &ordDetail)
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
