package usecases

import (
	"context"

	"github.com/masterraf21/reksti-ordering-backend/models"
)

type menuUsecase struct {
	menuRepo     models.MenuRepository
	menuTypeRepo models.MenuTypeRepository
}

// NewMenuUsecase will create new order usecase
func NewMenuUsecase(
	muc models.MenuRepository,
	mtc models.MenuTypeRepository,
) models.MenuUsecase {
	return &menuUsecase{
		menuRepo:     muc,
		menuTypeRepo: mtc,
	}
}

func (t *menuUsecase) GetAll() (res []models.MenuComp, err error) {
	menuRAW, err := t.menuRepo.GetAll()
	if err != nil {
		return
	}

	res = make([]models.MenuComp, len(menuRAW))
	for i, m := range menuRAW {
		res[i].MenuID = m.MenuID
		res[i].Name = m.Name
		res[i].Price = m.Price
		var typeTemp *models.MenuType
		typeTemp, err = t.menuTypeRepo.GetByID(m.MenuTypeID)
		if err != nil {
			return
		}
		res[i].MenuType = *typeTemp
		res[i].Ingredients = m.Ingredients
		res[i].MenuStatus = m.MenuStatus
	}
	return
}

func (t *menuUsecase) CreateMenu(ctx context.Context, order *models.Menu) (id uint32, err error) {
	id, err = t.menuRepo.Store(ctx, order)
	return
}

func (t *menuUsecase) DeleteMenu(ctx context.Context, id uint32) (res models.MenuComp, err error) {
	m, err := t.menuRepo.GetByID(id)
	err = t.menuRepo.DeleteByID(ctx, id)
	if err != nil {
		return
	}
	res.MenuID = m.MenuID
	res.Name = m.Name
	res.Price = m.Price
	typeTemp, err := t.menuTypeRepo.GetByID(m.MenuTypeID)
	if err != nil {
		return
	}
	res.MenuType = *typeTemp
	res.Ingredients = m.Ingredients
	res.MenuStatus = m.MenuStatus
	err = nil
	return
}

func (t *menuUsecase) GetByID(id uint32) (res *models.MenuComp, err error) {
	m, err := t.menuRepo.GetByID(id)
	if err != nil {
		return
	}

	if m == nil {
		return
	}

	typeTemp, err := t.menuTypeRepo.GetByID(m.MenuTypeID)
	if err != nil {
		return
	}
	menuComp := models.MenuComp{
		MenuID:      m.MenuID,
		Name:        m.Name,
		Price:       m.Price,
		MenuType:    *typeTemp,
		Ingredients: m.Ingredients,
		MenuStatus:  m.MenuStatus,
	}

	res = &menuComp

	return
}

func (t *menuUsecase) GetAllType() (res []models.MenuType, err error) {
	res, err = t.menuTypeRepo.GetAll()
	return
}

func (t *menuUsecase) CreateType(ctx context.Context, m *models.MenuType) (id uint32, err error) {
	id, err = t.menuTypeRepo.Store(context.Background(), m)
	return
}

func (t *menuUsecase) DeleteType(ctx context.Context, id uint32) (res *models.MenuType, err error) {
	res, err = t.menuTypeRepo.GetByID(id)
	if err != nil {
		return
	}
	err = t.menuTypeRepo.DeleteByID(ctx, id)
	return
}
