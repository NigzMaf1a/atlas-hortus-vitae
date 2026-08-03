package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	u User,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		u.SectorId,
		u.RoleId,
		u.UserName,
		u.Email,
		u.Password,
		u.AccStatus,
		u.RegType,
		u.Location,
	).Scan(&u.UserId)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	log.Println("User created successfully")
	return nil
}

func ReadUsers(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]User, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(
			&u.UserId,
			&u.SectorId,
			&u.RoleId,
			&u.UserName,
			&u.Email,
			&u.Password,
			&u.AccStatus,
			&u.RegType,
			&u.Location,
		); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate users: %w", err)
	}

	log.Println("Users fetched successfully")
	return users, nil
}

func UpdateAccStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {
	res, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)
	if err != nil {
		return fmt.Errorf("update account status: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	log.Println("Account status updated successfully")
	return nil
}

func UpdateUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	user User,
) error {
	res, err := db.ExecContext(
		ctx,
		query,
		user.Password,
		user.Location,
		user.UserId,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	log.Println("Account updated successfully")
	return nil
}
