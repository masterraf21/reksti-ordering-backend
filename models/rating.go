package models

// Rating model reting
type Rating struct {
	RatingID     uint32 `json:"rating_id"`
	MenuID       uint32 `json:"menu_id"`
	Score        int32  `json:"score"`
	Remarks      string `json:"remarks"`
	DateRecorded string `json:"date_recorded"`
	CustomerID   uint32 `json:"customer_id"`
}

// RatingRepository for repo
type RatingRepository interface {
	GetAll() ([]Rating, error)
	GetByID(ratingID uint32) (*Rating, error)
	GetByMenu(menuID uint32) ([]Rating, error)
	GetMenuScore(menuID uint32) (res float32, err error)
	UpdateByID(ratingID uint32, order *Rating) error
	DeleteByID(ratingID uint32) error
	Store(rating *Rating) (ratingID uint32, err error)
	BulkInsert(Rating []Rating) error
}
