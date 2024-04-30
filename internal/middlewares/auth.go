package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

const TokenSecret = "TGPTOfayPAqvUSRxRWhyyo4DsKwVxjQPJLa4Vim4u8E"
const NestfileToken = "nestfile-token"

type UserCtxKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenSecret := TokenSecret
		tokenString := extractTokenFromCookie(r)

		var ctx context.Context
		var claims *domain.AuthToken
		if tokenString != "" {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// You should implement your own logic to validate the token and return the appropriate key
				// For example, you could use a secret key or a public key
				return []byte(tokenSecret), nil
			})

			if err == nil || token.Valid {
				tokenClaims, err := transformMapClaimsToUserClaims(token.Claims)
				if err == nil {
					claims = tokenClaims
				}
			}

			ctx = context.WithValue(r.Context(), UserCtxKey{}, claims)
		} else {
			ctx = r.Context()
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie(NestfileToken)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func transformMapClaimsToUserClaims(claims jwt.Claims) (*domain.AuthToken, error) {
	if jwtClaims, ok := claims.(jwt.MapClaims); ok {
		sub, ok := jwtClaims["sub"].(float64)
		if !ok {
			return nil, fmt.Errorf("invalid 'sub' claim")
		}

		// username, ok := jwtClaims["username"].(string)
		// if !ok {
		// 	return nil, fmt.Errorf("invalid 'usename' claim")
		// }

		// Create a new instance of *UserClaims with the extracted values
		userClaims := &domain.AuthToken{
			Sub: int(sub),
		}

		// Add other custom claim fields here if needed
		return userClaims, nil
	}

	return nil, fmt.Errorf("failed to transform jwt.MapClaims to *UserClaims")
}
