package domain

import (
	"context"
	"time"

	"github.com/Brix101/nestfile/internal/util"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ResponseUser struct {
	User *User `json:"user"`
}

func (u *User) HashPwd() error {
	pwd, err := util.HashPwd(u.Password)
	if err != nil {
		return err
	}

	u.Password = pwd
	return nil
}

func (u User) CheckPwd(password string) bool {
	res := util.CheckPwd(password, u.Password)
	return res
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetAll(ctx context.Context) ([]User, error)

	// CreateOrUpdate(ctx context.Context, usr *User) error
	Update(ctx context.Context, usr *User) (*User, error)
	Create(ctx context.Context, usr *User) (*User, error)
	Delete(ctx context.Context, id int64) error
}
