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

// MenuRepository nn
type MenuRepository interface {
	GetAll() ([]Menu, error)
	GetById(menuID uint32) (*Menu, error)
	Update(ctx context.Context, menu *Menu) error
	DeleteByID(ctx context.Context, menuID uint32) error
}

// MenuTypeRepository repo
type MenuTypeRepository interface {
	GetAll() ([]MenuType, error)
	GetByID(menuTypeID uint32) (*MenuType, error)
	Update(ctx context.Context, mtype *MenuType) error
	DeleteByID(ctx context.Context, mtypeID uint32) error
}
