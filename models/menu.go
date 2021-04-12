package models

import "context"

// Menu for menu
type Menu struct {
	MenuID      uint32  `json:"menu_id"`
	Name        string  `json:"menu_name"`
	Price       float32 `json:"price"`
	MenuTypeID  uint32  `json:"menu_type_id"`
	Ingredients string  `json:"ingredients"`
	MenuStatus  bool    `json:"menu_status"`
}

// MenuType for menutype
type MenuType struct {
	MenuTypeID  uint32 `json:"menu_type_id"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}

// MenuComp :nodoc:
type MenuComp struct {
	MenuID      uint32   `json:"menu_id"`
	Name        string   `json:"menu_name"`
	Price       float32  `json:"price"`
	MenuType    MenuType `json:"menu_type"`
	Ingredients string   `json:"ingredients"`
	MenuStatus  bool     `json:"menu_status"`
}

// MenuRepository nn
type MenuRepository interface {
	GetAll() ([]Menu, error)
	GetByID(menuID uint32) (*Menu, error)
	UpdateByID(ctx context.Context, menuID uint32, order *Menu) error
	DeleteByID(ctx context.Context, menuID uint32) error
	Store(ctx context.Context, ord *Menu) (menuID uint32, err error)
	BulkInsert(ctx context.Context, Menu []Menu) error
}

// MenuTypeRepository repo
type MenuTypeRepository interface {
	GetAll() ([]MenuType, error)
	GetByID(menuTypeID uint32) (*MenuType, error)
	UpdateByID(ctx context.Context, menuTypeID uint32, order *MenuType) error
	DeleteByID(ctx context.Context, mtypeID uint32) error
	Store(ctx context.Context, ord *MenuType) (menuTypeID uint32, err error)
	BulkInsert(ctx context.Context, MenuType []MenuType) error
}

// MenuUsecase :nododc:
type MenuUsecase interface {
	GetAll() (res []MenuComp, err error)
	CreateMenu(ctx context.Context, order *Menu) (id uint32, err error)
	DeleteMenu(ctx context.Context, id uint32) (res MenuComp, err error)
	GetByID(id uint32) (res *MenuComp, err error)
	GetAllType() (res []MenuType, err error)
	CreateType(ctx context.Context, m *MenuType) (id uint32, err error)
	DeleteType(ctx context.Context, id uint32) (res *MenuType, err error)
}
