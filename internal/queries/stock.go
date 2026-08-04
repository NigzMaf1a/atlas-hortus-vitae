package queries

var StockQueries = struct {
	CreateStock      string
	ReadStock        string
	ReadOutletStock  string
	UpdateStockQty   string
	UpdateStockPrice string
}{
	CreateStock: `
		INSERT INTO stock (
			outlet_id,
			stock_name,
			stock_qty,
			cost_per_kg
		)
		VALUES ($1, $2, $3, $4)
		RETURNING stock_id
	`,

	ReadStock: `
		SELECT
			stock_id,
			outlet_id,
			stock_name,
			stock_qty,
			cost_per_kg
		FROM stock
		ORDER BY stock_id
	`,

	ReadOutletStock: `
		SELECT
			stock_id,
			outlet_id,
			stock_name,
			stock_qty,
			cost_per_kg
		FROM stock
		WHERE outlet_id = $1
		ORDER BY stock_id
	`,

	UpdateStockQty: `
		UPDATE stock
		SET stock_qty = $1
		WHERE stock_id = $2
	`,

	UpdateStockPrice: `
		UPDATE stock
		SET cost_per_kg = $1
		WHERE stock_id = $2
	`,
}
