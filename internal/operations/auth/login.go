package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func Login(
	ctx context.Context,
	db *sql.DB,
	query string,
	l LoginCred,
) (User, error) {
	var u User
	log.Printf("Email: %s", l.Email)
	log.Printf("Password: %s", l.Password)

	err := db.QueryRowContext(
		ctx,
		query,
		l.Email,
		l.Password,
	).Scan(
		&u.UserId,
		&u.SectorId,
		&u.RoleId,
		&u.UserName,
		&u.Email,
		&u.Password,
		&u.AccStatus,
		&u.RegType,
		&u.Location,
	)

	if err != nil {
		log.Printf("Login error: %v", err)

		if err == sql.ErrNoRows {
			return u, sql.ErrNoRows
		}

		return u, fmt.Errorf("login user: %w", err)
	}

	log.Println("User authenticated successfully")

	return u, nil
}
