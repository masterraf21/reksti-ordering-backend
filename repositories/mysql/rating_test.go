package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type ratingRepoTestSuite struct {
	suite.Suite
	Reader     *sql.DB
	Writer     *sql.DB
	Repo       models.RatingRepository
	CustomerID uint32
	MenuID     uint32
}

func TestRatingRepository(t *testing.T) {
	suite.Run(t, new(ratingRepoTestSuite))
}

func (s *ratingRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewRatingRepo(reader, writer)
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
}

func (s *ratingRepoTestSuite) TearDownSuite() {
	querys := []string{
		"DELETE FROM rating;",
		"DELETE FROM customer;",
		"DELETE FROM menu;",
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

func (s *ratingRepoTestSuite) TearDownTest() {
	querys := []string{
		"DELETE FROM rating;",
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

func (s *ratingRepoTestSuite) TestInserts() {
	s.Run("Store Rating", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		id, err := s.Repo.Store(&rating)
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
		s.Equal(id, result.RatingID)
		s.Equal(rating.MenuID, result.MenuID)
		s.Equal(rating.Score, result.Score)
		s.Equal(rating.Remarks, result.Remarks)
		s.Equal(rating.CustomerID, result.CustomerID)
	})
}

func (s *ratingRepoTestSuite) TestGet() {
	s.Run("Get Rating By Id", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		id, err := s.Repo.Store(&rating)
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
		s.Equal(id, result.RatingID)
		s.Equal(rating.MenuID, result.MenuID)
		s.Equal(rating.Score, result.Score)
		s.Equal(rating.Remarks, result.Remarks)
		s.Equal(rating.CustomerID, result.CustomerID)
	})

	s.Run("Get Rating by Menu", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		_, err := s.Repo.Store(&rating)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		results, err := s.Repo.GetByMenu(rating.MenuID)
		result := results[0]
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(rating.MenuID, result.MenuID)
		s.Equal(rating.Score, result.Score)
		s.Equal(rating.Remarks, result.Remarks)
		s.Equal(rating.CustomerID, result.CustomerID)
	})

	s.Run("Get Menu Score", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		id, err := s.Repo.Store(&rating)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		results, err := s.Repo.GetMenuScore(s.MenuID)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(float32(rating.Score), results)
	})
}

func (s *ratingRepoTestSuite) TestGet2() {
	s.Run("Get All Rating", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(&rating)
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

func (s *ratingRepoTestSuite) TestUpdate() {
	s.Run("Update Rating", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		id, err := s.Repo.Store(&rating)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)
		rating2 := models.Rating{
			MenuID:       s.MenuID,
			Score:        5,
			Remarks:      "Tested by Tape",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}

		err = s.Repo.UpdateByID(id, &rating2)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)

		result, err := s.Repo.GetByID(id)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Equal(rating2.Score, result.Score)
		s.Equal(rating2.Remarks, result.Remarks)
	})
}

func (s *ratingRepoTestSuite) TestDelete() {
	s.Run("Delete Order By ID", func() {
		rating := models.Rating{
			MenuID:       s.MenuID,
			Score:        7,
			Remarks:      "Halooo",
			DateRecorded: "2021-04-14",
			CustomerID:   s.CustomerID,
		}
		id, err := s.Repo.Store(&rating)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		err = s.Repo.DeleteByID(id)
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
