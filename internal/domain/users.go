package domain

import (
	"context"
	"time"

	"github.com/Brix101/nestfile/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u User) CheckPwd(password string) bool {
	res := util.CheckPwd(password, u.Password)
	return res
}

func (u *User) HashPwd() error {
	pwd, err := util.HashPwd(u.Password)
	if err != nil {
		return err
	}

	u.Password = pwd
	return nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Sub      int    `json:"sub"`
}

// TODO move this constant into a config
const TokenSecret = "TGPTOfayPAqvUSRxRWhyyo4DsKwVxjQPJLa4Vim4u8E"

func (u User) GenerateClaims() (string, error) {
	tokenSecret := TokenSecret
	claims := UserClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
		u.Username,
		int(u.ID),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return t, nil
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
