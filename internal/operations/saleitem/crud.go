package saleitem

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateSaleItem(
	ctx context.Context,
	db *sql.DB,
	query string,
	s SaleItem,
) (SaleItem, error) {

	err := db.QueryRowContext(
		ctx,
		query,
		s.SaleID,
		s.ProductID,
		s.ProductPrice,
		s.SaleQty,
		s.SaleTotal,
	).Scan(&s.SaleItemID)

	if err != nil {
		return s, fmt.Errorf("create sale item: %w", err)
	}

	fmt.Println("Sale item created successfully")

	return s, nil
}

func ReadSaleItems(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]SaleItem, error) {

	saleItems := []SaleItem{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read sale items: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s SaleItem

		err := rows.Scan(
			&s.SaleItemID,
			&s.SaleID,
			&s.ProductID,
			&s.ProductPrice,
			&s.SaleQty,
			&s.SaleTotal,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale item: %w", err)
		}

		saleItems = append(saleItems, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sale items: %w", err)
	}

	fmt.Println("Sale items fetched successfully")

	return saleItems, nil
}

func ReadSaleItemsBySale(
	ctx context.Context,
	db *sql.DB,
	query string,
	saleID int64,
) ([]SaleItem, error) {

	saleItems := []SaleItem{}

	rows, err := db.QueryContext(
		ctx,
		query,
		saleID,
	)

	if err != nil {
		return nil, fmt.Errorf("read sale items by sale: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s SaleItem

		err := rows.Scan(
			&s.SaleItemID,
			&s.SaleID,
			&s.ProductID,
			&s.ProductPrice,
			&s.SaleQty,
			&s.SaleTotal,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale item: %w", err)
		}

		saleItems = append(saleItems, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sale items by sale: %w", err)
	}

	fmt.Println("Sale items for sale fetched successfully")

	return saleItems, nil
}

func ReadSaleItemsByProduct(
	ctx context.Context,
	db *sql.DB,
	query string,
	productID int64,
) ([]SaleItem, error) {

	saleItems := []SaleItem{}

	rows, err := db.QueryContext(
		ctx,
		query,
		productID,
	)

	if err != nil {
		return nil, fmt.Errorf("read sale items by product: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var s SaleItem

		err := rows.Scan(
			&s.SaleItemID,
			&s.SaleID,
			&s.ProductID,
			&s.ProductPrice,
			&s.SaleQty,
			&s.SaleTotal,
		)

		if err != nil {
			return nil, fmt.Errorf("scan sale item: %w", err)
		}

		saleItems = append(saleItems, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sale items by product: %w", err)
	}

	fmt.Println("Sale items for product fetched successfully")

	return saleItems, nil
}
