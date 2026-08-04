package outlets

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateOutlet(
	ctx context.Context,
	db *sql.DB,
	query string,
	o Outlet,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		&o.Name,
		&o.Location,
		&o.Networth,
		&o.Networth,
		&o.Open,
	).Scan(&o.OutletID)

	if err != nil {
		fmt.Println("An error occurred while querying the database")
		return fmt.Errorf("create outlet: %w", err)
	}

	fmt.Println("Outlet created successfully")

	return nil
}

func ReadOutlets(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Outlet, error) {
	outlets := []Outlet{}

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		fmt.Println("An error occurred while querying the database")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var o Outlet

		err := rows.Scan(
			&o.OutletID,
			&o.Name,
			&o.Location,
			&o.Networth,
			&o.Open,
		)

		if err != nil {
			fmt.Println("An error occurred while scanning a record")
			return nil, err
		}

		outlets = append(outlets, o)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate outlets: %w", err)
	}

	fmt.Println("Outlets fetched successfully")

	return outlets, nil
}

func ReadOutlet(
	ctx context.Context,
	db *sql.DB,
	query string,
	id int64,
) (Outlet, error) {
	var o Outlet

	err := db.QueryRowContext(
		ctx, query, id,
	).Scan(
		&o.OutletID,
		&o.Name,
		&o.Location,
		&o.Networth,
		&o.Open,
	)

	if err != nil {
		fmt.Println("An error occurred while querying the database")
		return o, err
	}

	fmt.Println("Outlet fetched successfully")

	return o, err
}

func UpdateNetworth(
	ctx context.Context,
	db *sql.DB,
	query string,
	shs float64,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx, query, shs, id,
	)

	if err != nil {
		fmt.Println("An error occurred while querying the DB")
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Newtworth update successful")

	return nil
}

func UpdateOutletStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx, query, status, id,
	)

	if err != nil {
		fmt.Println("An error occurred while querying the DB")
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if aff == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Outlet status update successful")

	return nil
}
