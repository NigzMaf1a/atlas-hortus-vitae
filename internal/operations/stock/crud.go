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
) (Stock, error) {

	err := db.QueryRowContext(
		ctx,
		query,
		s.OutletID,
		s.StockName,
		s.StockQty,
		s.CostPerKg,
	).Scan(&s.StockID)

	if err != nil {
		return s, fmt.Errorf("create stock: %w", err)
	}

	fmt.Println("Stock created successfully")

	return s, nil
}

func ReadStock(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Stock, error) {

	stocks := []Stock{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read stock: %w", err)
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
			return nil, fmt.Errorf("scan stock: %w", err)
		}

		stocks = append(stocks, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate stock: %w", err)
	}

	fmt.Println("Stock fetched successfully")

	return stocks, nil
}

func ReadOutletStock(
	ctx context.Context,
	db *sql.DB,
	query string,
	outletID int64,
) ([]Stock, error) {

	stocks := []Stock{}

	rows, err := db.QueryContext(
		ctx,
		query,
		outletID,
	)

	if err != nil {
		return nil, fmt.Errorf("read outlet stock: %w", err)
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
			return nil, fmt.Errorf("scan outlet stock: %w", err)
		}

		stocks = append(stocks, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate outlet stock: %w", err)
	}

	fmt.Println("Outlet stock fetched successfully")

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
		ctx,
		query,
		qty,
		id,
	)

	if err != nil {
		return fmt.Errorf("update stock quantity: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
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
		ctx,
		query,
		cost,
		id,
	)

	if err != nil {
		return fmt.Errorf("update stock cost: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Cost update successful")

	return nil
}
