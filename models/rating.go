package models

import "context"

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
	GetById(ratingID uint32) (*Rating, error)
	Update(ctx context.Context, rating *Rating) error
	DeleteById(ratingID uint32) error
	Store(rating *Rating) error
}
