package payments

import (
	"context"
	"database/sql"
	"fmt"
)

func CreatePayment(
	ctx context.Context,
	db *sql.DB,
	query string,
	p Payment,
) (Payment, error) {

	err := db.QueryRowContext(
		ctx,
		query,
		p.UserId,
		p.SaleID,
		p.SalePrice,
		p.PaymentCode,
		p.PaymentStatus,
	).Scan(&p.PaymentID)

	if err != nil {
		return p, fmt.Errorf("create payment: %w", err)
	}

	fmt.Println("Payment created successfully")

	return p, nil
}

func UpdatePaymentStatus(
	ctx context.Context,
	db *sql.DB,
	query string,
	status string,
	id int64,
) error {

	result, err := db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		return fmt.Errorf("update payment status: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("retrieve affected rows: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Payment status updated successfully")

	return nil
}

func ReadPayments(
	ctx context.Context,
	db *sql.DB,
	query string,
) ([]Payment, error) {

	payments := []Payment{}

	rows, err := db.QueryContext(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("read payments: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var p Payment

		err := rows.Scan(
			&p.PaymentID,
			&p.UserId,
			&p.PaymentDate,
			&p.SaleID,
			&p.SalePrice,
			&p.PaymentCode,
			&p.PaymentStatus,
		)

		if err != nil {
			return nil, fmt.Errorf("scan payment: %w", err)
		}

		payments = append(payments, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate payments: %w", err)
	}

	fmt.Println("Payments fetched successfully")

	return payments, nil
}

func ReadPaymentsByUser(
	ctx context.Context,
	db *sql.DB,
	query string,
	userID int64,
) ([]Payment, error) {

	payments := []Payment{}

	rows, err := db.QueryContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		return nil, fmt.Errorf("read payments by user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {

		var p Payment

		err := rows.Scan(
			&p.PaymentID,
			&p.UserId,
			&p.PaymentDate,
			&p.SaleID,
			&p.SalePrice,
			&p.PaymentCode,
			&p.PaymentStatus,
		)

		if err != nil {
			return nil, fmt.Errorf("scan payment: %w", err)
		}

		payments = append(payments, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate user payments: %w", err)
	}

	fmt.Println("User payments fetched successfully")

	return payments, nil
}
