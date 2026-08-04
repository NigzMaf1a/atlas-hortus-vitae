package stock

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateStock(
	ctx context.Context,
	db *sql.DB,
	query string,
	s Stock,
) error {
	err := db.QueryRowContext(
		ctx,
		query,
		&s.OutletID,
		&s.StockName,
		&s.StockQty,
		&s.CostPerKg,
	).Scan(&s.StockID)

	if err != nil {
		fmt.Println("An error occurred while querying the database")
		return fmt.Errorf("create stock: %w", err)
	}

	fmt.Println("Stock created successfully")

	return nil
}

func ReadOutletStock(
	ctx context.Context,
	db *sql.DB,
	query string,
	outlet_id int64,
) ([]Stock, error) {
	stocks := []Stock{}

	rows, err := db.QueryContext(
		ctx, query, outlet_id,
	)

	if err != nil {
		fmt.Println("An error occurred while querying the database")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s Stock

		err := rows.Scan(
			&s.StockID,
			&s.OutletID,
			&s.StockName,
			&s.StockQty,
			&s.CostPerKg,
		)

		if err != nil {
			fmt.Println("An error occurred while scanning a record")
			return nil, err
		}

		stocks = append(stocks, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate stock: %w", err)
	}

	fmt.Println("Stock fetched successfully")

	return stocks, nil
}

func UpdateStockQty(
	ctx context.Context,
	db *sql.DB,
	query string,
	qty float64,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx, query, qty, id,
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

	fmt.Println("Quantity update successful")

	return nil
}

func UpdateStockPrice(
	ctx context.Context,
	db *sql.DB,
	query string,
	cost float64,
	id int64,
) error {
	result, err := db.ExecContext(
		ctx, query, cost, id,
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

	fmt.Println("Cost update successful")

	return nil
}
