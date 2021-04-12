package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type menuRepositoryTestSuite struct {
	suite.Suite
	Reader     *sql.DB
	Writer     *sql.DB
	Repo       models.MenuRepository
	MenuTypeID uint32
}

func TestMenuRepository(t *testing.T) {
	suite.Run(t, new(menuRepositoryTestSuite))
}

func (s *menuRepositoryTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewMenuRepo(reader, writer)
	menuTypeRepo := NewMenuTypeRepo(reader, writer)

	mt := models.MenuType{
		TypeName:    "Siomay",
		Description: "Ueank",
	}
	id, err := menuTypeRepo.Store(context.TODO(), &mt)
	if err != nil {
		panic(err)
	}
	s.MenuTypeID = id
}

func (s *menuRepositoryTestSuite) TearDownSuite() {
	querys := []string{
		"DELETE FROM menu;",
		"DELETE FROM menu_type;",
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

func (s *menuRepositoryTestSuite) TearDownTest() {
	querys := []string{
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

func (s *menuRepositoryTestSuite) TestStore() {
	s.Run("Store Menu", func() {
		menu := models.Menu{
			MenuTypeID:  s.MenuTypeID,
			Name:        "Mario",
			Price:       100000,
			Ingredients: "Luigi Goreng",
			MenuStatus:  true,
		}
		id, err := s.Repo.Store(context.TODO(), &menu)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)
	})
}

func (s *menuRepositoryTestSuite) TestGet() {
	s.Run("Get Menu By ID", func() {
		menu := models.Menu{
			MenuTypeID:  s.MenuTypeID,
			Name:        "Mario",
			Price:       100000,
			Ingredients: "Luigi Goreng",
			MenuStatus:  true,
		}
		id, err := s.Repo.Store(context.TODO(), &menu)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)

		result, err := s.Repo.GetByID(id)
		s.Equal(id, result.MenuID)
		s.Equal(menu.MenuTypeID, result.MenuTypeID)
		s.Equal(menu.Name, result.Name)
		s.Equal(menu.Price, result.Price)
		s.Equal(menu.Ingredients, result.Ingredients)
	})
}

func (s *menuRepositoryTestSuite) TestGet2() {
	s.Run("Get All Menu", func() {
		menu := models.Menu{
			MenuTypeID:  s.MenuTypeID,
			Name:        "Mario",
			Price:       100000,
			Ingredients: "Luigi Goreng",
			MenuStatus:  true,
		}
		for i := 0; i < 5; i++ {
			id, err := s.Repo.Store(context.TODO(), &menu)
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

func (s *menuRepositoryTestSuite) TestDelete() {
	s.Run("Get Menu By ID", func() {
		menu := models.Menu{
			MenuTypeID:  s.MenuTypeID,
			Name:        "Mario",
			Price:       100000,
			Ingredients: "Luigi Goreng",
			MenuStatus:  true,
		}
		id, err := s.Repo.Store(context.TODO(), &menu)
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
		s.Assert().Nil(result)
	})
}
