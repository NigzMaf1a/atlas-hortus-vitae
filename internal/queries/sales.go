package queries

var SaleQueries = struct {
	CreateSale       string
	ReadSales        string
	ReadSalesByUser  string
	ReadSalesByOutlet string
	ReadSalesByDate  string
}{
	CreateSale: `
		INSERT INTO sales (
			user_id,
			outlet_id,
			sale_date,
			sale_total,
			sale_discount,
			sale_price
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING sale_id
	`,

	ReadSales: `
		SELECT
			sale_id,
			user_id,
			outlet_id,
			sale_date,
			sale_total,
			sale_discount,
			sale_price
		FROM sales
		ORDER BY sale_date DESC
	`,

	ReadSalesByUser: `
		SELECT
			sale_id,
			user_id,
			outlet_id,
			sale_date,
			sale_total,
			sale_discount,
			sale_price
		FROM sales
		WHERE user_id = $1
		ORDER BY sale_date DESC
	`,

	ReadSalesByOutlet: `
		SELECT
			sale_id,
			user_id,
			outlet_id,
			sale_date,
			sale_total,
			sale_discount,
			sale_price
		FROM sales
		WHERE outlet_id = $1
		ORDER BY sale_date DESC
	`,

	ReadSalesByDate: `
		SELECT
			sale_id,
			user_id,
			outlet_id,
			sale_date,
			sale_total,
			sale_discount,
			sale_price
		FROM sales
		WHERE sale_date >= $1
		AND sale_date < $1 + INTERVAL '1 day'
		ORDER BY sale_date DESC
	`,
}