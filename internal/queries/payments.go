package queries

var PaymentQueries = struct {
	CreatePayment       string
	UpdatePaymentStatus string
	ReadPayments        string
	ReadPaymentsByUser  string
}{
	CreatePayment: `
		INSERT INTO payments (
			user_id,
			sale_id,
			sale_price,
			payment_code,
			payment_status
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING payment_id
	`,

	UpdatePaymentStatus: `
		UPDATE payments
		SET payment_status = $1
		WHERE payment_id = $2
	`,

	ReadPayments: `
		SELECT
			payment_id,
			user_id,
			payment_date,
			sale_id,
			sale_price,
			payment_code,
			payment_status
		FROM payments
		ORDER BY payment_date DESC
	`,

	ReadPaymentsByUser: `
		SELECT
			payment_id,
			user_id,
			payment_date,
			sale_id,
			sale_price,
			payment_code,
			payment_status
		FROM payments
		WHERE user_id = $1
		ORDER BY payment_date DESC
	`,
}
