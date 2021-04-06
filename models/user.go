package models

import "context"

// User object
type User struct {
	UserID   uint32 `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email_address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRepository for user tables
type UserRepository interface {
	GetAll() ([]User, error)
	GetById(UserID uint32) (*User, error)
	Update(ctx context.Context, user *User) error
	DeleteById(ctx context.Context, UserID uint32) error
	Store(ctx context.Context, user *User) error
}
