package products

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateProduct(
	ctx context.Context,
	db *sql.DB,
	query string,
	p Product,
) (Product, error) {

	err := db.QueryRowContext(
		ctx,
		query,
		p.OutletID,
		p.ProductName,
		p.ProductPrice,
		p.Available,
	).Scan(&p.ProductID)

	if err != nil {
		return p, fmt.Errorf("create product: %w", err)
	}

	fmt.Println("Product created successfully")

	return p, nil
}

func ReadProducts(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Product, error) {

	products := []Product{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read products: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var p Product

		err := rows.Scan(
			&p.ProductID,
			&p.OutletID,
			&p.ProductName,
			&p.ProductPrice,
			&p.Available,
		)

		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate products: %w", err)
	}

	fmt.Println("Products fetched successfully")

	return products, nil
}

func ReadProductsByOutlet(
	ctx context.Context,
	db *sql.DB,
	query string,
	outletID int64,
) ([]Product, error) {

	products := []Product{}

	rows, err := db.QueryContext(
		ctx,
		query,
		outletID,
	)

	if err != nil {
		return nil, fmt.Errorf("read products by outlet: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var p Product

		err := rows.Scan(
			&p.ProductID,
			&p.OutletID,
			&p.ProductName,
			&p.ProductPrice,
			&p.Available,
		)

		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate outlet products: %w", err)
	}

	fmt.Println("Outlet products fetched successfully")

	return products, nil
}

func ReadProductsByAvailable(
	ctx context.Context,
	db *sql.DB,
	query string,
	available string,
) ([]Product, error) {

	products := []Product{}

	rows, err := db.QueryContext(
		ctx,
		query,
		available,
	)

	if err != nil {
		return nil, fmt.Errorf("read available products: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var p Product

		err := rows.Scan(
			&p.ProductID,
			&p.OutletID,
			&p.ProductName,
			&p.ProductPrice,
			&p.Available,
		)

		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate available products: %w", err)
	}

	fmt.Println("Available products fetched successfully")

	return products, nil
}

func UpdateProductPrice(
	ctx context.Context,
	db *sql.DB,
	query string,
	price int64,
	id int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		price,
		id,
	)

	if err != nil {
		return fmt.Errorf("update product price: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Product price updated successfully")

	return nil
}

func UpdateProductAvailable(
	ctx context.Context,
	db *sql.DB,
	query string,
	available string,
	id int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		available,
		id,
	)

	if err != nil {
		return fmt.Errorf("update product availability: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Product availability updated successfully")

	return nil
}
