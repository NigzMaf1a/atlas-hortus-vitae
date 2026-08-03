package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/auth"
)

var JWTSecret = []byte("your-super-secret-key")

func GenerateJWT(user auth.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    user.UserId,
		"sector_id":  user.SectorId,
		"role_id":    user.RoleId,
		"user_name":  user.UserName,
		"email":      user.Email,
		"acc_status": user.AccStatus,
		"reg_type":   user.RegType,
		"location":   user.Location,

		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JWTSecret)
}
