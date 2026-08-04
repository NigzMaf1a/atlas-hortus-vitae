package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("your-super-secret-key")

type contextKey string

const (
	UserIDKey    contextKey = "user_id"
	SectorIDKey  contextKey = "sector_id"
	RoleIDKey    contextKey = "role_id"
	UserNameKey  contextKey = "user_name"
	EmailKey     contextKey = "email"
	AccStatusKey contextKey = "acc_status"
	RegTypeKey   contextKey = "reg_type"
	LocationKey  contextKey = "location"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization required", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				return JWTSecret, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, UserIDKey, claims["user_id"])
		ctx = context.WithValue(ctx, SectorIDKey, claims["sector_id"])
		ctx = context.WithValue(ctx, RoleIDKey, claims["role_id"])
		ctx = context.WithValue(ctx, UserNameKey, claims["user_name"])
		ctx = context.WithValue(ctx, EmailKey, claims["email"])
		ctx = context.WithValue(ctx, AccStatusKey, claims["acc_status"])
		ctx = context.WithValue(ctx, RegTypeKey, claims["reg_type"])
		ctx = context.WithValue(ctx, LocationKey, claims["location"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
