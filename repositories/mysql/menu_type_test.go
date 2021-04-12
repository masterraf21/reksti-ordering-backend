package mysql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type menuTypeRepoTestSuite struct {
	suite.Suite
	Reader *sql.DB
	Writer *sql.DB
	Repo   models.MenuTypeRepository
}

func TestMenuTypeRepository(t *testing.T) {
	suite.Run(t, new(menuTypeRepoTestSuite))
}

func (s *menuTypeRepoTestSuite) SetupSuite() {
	reader, writer := configureMySQL()
	s.Reader = reader
	s.Writer = writer
	s.Repo = NewMenuTypeRepo(reader, writer)
}

func (s *menuTypeRepoTestSuite) TearDownSuite() {
	querys := []string{
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

func (s *menuTypeRepoTestSuite) TearDownTest() {
	querys := []string{
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

func (s *menuTypeRepoTestSuite) TestStore() {
	s.Run("Store Menu Type", func() {
		mt := models.MenuType{
			TypeName:    "Siomay",
			Description: "Ueank",
		}
		id, err := s.Repo.Store(context.TODO(), &mt)
		if err != nil {
			panic(err)
		}
		s.Require().NoError(err)
		s.Assert().NotEmpty(id)
	})
}
