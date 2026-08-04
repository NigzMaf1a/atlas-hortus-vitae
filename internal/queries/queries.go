package queries

var ProductQueries = struct {
	CreateProduct          string
	ReadProducts           string
	ReadProductsByOutlet   string
	ReadProductsByAvailable string
	UpdateProductPrice     string
	UpdateProductAvailable string
}{
	CreateProduct: `
		INSERT INTO products (
			outlet_id,
			product_name,
			product_price,
			available
		)
		VALUES ($1, $2, $3, $4)
		RETURNING product_id
	`,

	ReadProducts: `
		SELECT
			product_id,
			outlet_id,
			product_name,
			product_price,
			available
		FROM products
		ORDER BY product_id
	`,

	ReadProductsByOutlet: `
		SELECT
			product_id,
			outlet_id,
			product_name,
			product_price,
			available
		FROM products
		WHERE outlet_id = $1
		ORDER BY product_id
	`,

	ReadProductsByAvailable: `
		SELECT
			product_id,
			outlet_id,
			product_name,
			product_price,
			available
		FROM products
		WHERE available = $1
		ORDER BY product_id
	`,

	UpdateProductPrice: `
		UPDATE products
		SET product_price = $1
		WHERE product_id = $2
	`,

	UpdateProductAvailable: `
		UPDATE products
		SET available = $1
		WHERE product_id = $2
	`,
}