package queries

var SaleItemQueries = struct {
	CreateSaleItem         string
	ReadSaleItems          string
	ReadSaleItemsBySale    string
	ReadSaleItemsByProduct string
}{
	CreateSaleItem: `
		INSERT INTO sale_items (
			sale_id,
			product_id,
			product_price,
			sale_qty,
			sale_total
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING sale_item_id
	`,

	ReadSaleItems: `
		SELECT
			sale_item_id,
			sale_id,
			product_id,
			product_price,
			sale_qty,
			sale_total
		FROM sale_items
		ORDER BY sale_item_id
	`,

	ReadSaleItemsBySale: `
		SELECT
			sale_item_id,
			sale_id,
			product_id,
			product_price,
			sale_qty,
			sale_total
		FROM sale_items
		WHERE sale_id = $1
		ORDER BY sale_item_id
	`,

	ReadSaleItemsByProduct: `
		SELECT
			sale_item_id,
			sale_id,
			product_id,
			product_price,
			sale_qty,
			sale_total
		FROM sale_items
		WHERE product_id = $1
		ORDER BY sale_item_id
	`,
}
