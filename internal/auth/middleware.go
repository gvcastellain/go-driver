package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer")

		claims := new(Claims)

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(rw, err.Error(), http.StatusProxyAuthRequired)
				return
			}

			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(rw, err.Error(), http.StatusProxyAuthRequired)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "user_name", claims.UserName)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
