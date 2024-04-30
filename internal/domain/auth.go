package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthToken struct {
	jwt.RegisteredClaims
	Sub int `json:"sub"`
}

// TODO move this constant into a config
const TokenSecret = "TGPTOfayPAqvUSRxRWhyyo4DsKwVxjQPJLa4Vim4u8E"

func (u User) GenerateClaims() (string, error) {
	tokenSecret := TokenSecret
	claims := AuthToken{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "Nestfile",
		},
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
