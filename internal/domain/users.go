package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRepository represents the user's repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetAll(ctx context.Context) ([]User, error)

	// CreateOrUpdate(ctx context.Context, usr *User) error
	Update(ctx context.Context, usr *User) (*User, error)
	Create(ctx context.Context, usr *User) (*User, error)
	Delete(ctx context.Context, id int64) error
}
