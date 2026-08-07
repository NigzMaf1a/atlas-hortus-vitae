package sales

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/saleitem"
)

func CreateSale(
	ctx context.Context,
	db *sql.DB,
	saleQuery string,
	saleItemQuery string,
	s Sale,
	items []saleitem.SaleItem,
) (Sale, []saleitem.SaleItem, error) {

	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return s, nil, fmt.Errorf("begin sale transaction: %w", err)
	}

	// Rollback is safe even if Commit succeeds.
	defer tx.Rollback()

	// Create sale
	err = tx.QueryRowContext(
		ctx,
		saleQuery,
		s.UserId,
		s.OutletID,
		s.SaleDate,
		s.SaleStatus,
		s.SaleTotal,
		s.SaleDiscount,
		s.SalePrice,
	).Scan(&s.SaleID)

	if err != nil {
		return s, nil, fmt.Errorf("create sale: %w", err)
	}

	// Create sale items
	createdItems := make([]saleitem.SaleItem, 0, len(items))

	for _, item := range items {

		item.SaleID = s.SaleID

		err := tx.QueryRowContext(
			ctx,
			saleItemQuery,
			item.SaleID,
			item.ProductID,
			item.ProductPrice,
			item.SaleQty,
			item.SaleTotal,
		).Scan(&item.SaleItemID)

		if err != nil {
			return s, nil, fmt.Errorf(
				"create sale item for sale %d: %w",
				s.SaleID,
				err,
			)
		}

		createdItems = append(createdItems, item)
	}

	// Commit everything
	if err := tx.Commit(); err != nil {
		return s, nil, fmt.Errorf("commit sale: %w", err)
	}

	fmt.Println("Sale and sale items created successfully")

	return s, createdItems, nil
}

func ReadSales(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Sale, error) {

	sales := []Sale{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read sales: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s Sale

		err := rows.Scan(
			&s.SaleID,
			&s.UserId,
			&s.OutletID,
			&s.SaleDate,
			&s.SaleStatus,
			&s.SaleTotal,
			&s.SaleDiscount,
			&s.SalePrice,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale: %w", err)
		}

		sales = append(sales, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sales: %w", err)
	}

	fmt.Println("Sales fetched successfully")

	return sales, nil
}

func ReadSalesByUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	userID int64,
) ([]Sale, error) {

	sales := []Sale{}

	rows, err := db.QueryContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, fmt.Errorf("read sales by user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s Sale

		err := rows.Scan(
			&s.SaleID,
			&s.UserId,
			&s.OutletID,
			&s.SaleDate,
			&s.SaleStatus,
			&s.SaleTotal,
			&s.SaleDiscount,
			&s.SalePrice,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale: %w", err)
		}

		sales = append(sales, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate user sales: %w", err)
	}

	fmt.Println("User sales fetched successfully")

	return sales, nil
}

func ReadSalesByOutlet(
	ctx context.Context,
	db *sql.DB,
	query string,
	outletID int64,
) ([]Sale, error) {

	sales := []Sale{}

	rows, err := db.QueryContext(
		ctx,
		query,
		outletID,
	)

	if err != nil {
		return nil, fmt.Errorf("read sales by outlet: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s Sale

		err := rows.Scan(
			&s.SaleID,
			&s.UserId,
			&s.OutletID,
			&s.SaleDate,
			&s.SaleStatus,
			&s.SaleTotal,
			&s.SaleDiscount,
			&s.SalePrice,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale: %w", err)
		}

		sales = append(sales, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate outlet sales: %w", err)
	}

	fmt.Println("Outlet sales fetched successfully")

	return sales, nil
}

func ReadSalesByDate(
	ctx context.Context,
	db *sql.DB,
	query string,
	date time.Time,
) ([]Sale, error) {

	sales := []Sale{}

	rows, err := db.QueryContext(
		ctx,
		query,
		date,
	)

	if err != nil {
		return nil, fmt.Errorf("read sales by date: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s Sale

		err := rows.Scan(
			&s.SaleID,
			&s.UserId,
			&s.OutletID,
			&s.SaleDate,
			&s.SaleStatus,
			&s.SaleTotal,
			&s.SaleDiscount,
			&s.SalePrice,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale: %w", err)
		}

		sales = append(sales, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sales by date: %w", err)
	}

	fmt.Println("Sales for date fetched successfully")

	return sales, nil
}

func UpdateSaleStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	saleID int64,
	status string,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		status,
		saleID,
	)

	if err != nil {
		return fmt.Errorf("update sale status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no sale found with id %d", saleID)
	}

	fmt.Printf(
		"Sale %d status updated successfully to %s\n",
		saleID,
		status,
	)

	return nil
}
