package queries

var OutletQueries = struct {
	CreateOutlet       string
	ReadOutlets        string
	ReadOutlet         string
	UpdateNetworth     string
	UpdateOutletStatus string
}{
	CreateOutlet: `
		INSERT INTO outlets (
			name,
			location,
			networth,
			open
		)
		VALUES ($1, $2, $3, $4)
		RETURNING outlet_id
	`,

	ReadOutlets: `
		SELECT
			outlet_id,
			name,
			location,
			networth,
			open
		FROM outlets
		ORDER BY outlet_id
	`,

	ReadOutlet: `
		SELECT
			outlet_id,
			name,
			location,
			networth,
			open
		FROM outlets
		WHERE outlet_id = $1
	`,

	UpdateNetworth: `
		UPDATE outlets
		SET networth = networth + $1
		WHERE outlet_id = $2
	`,

	UpdateOutletStatus: `
		UPDATE outlets
		SET open = $1
		WHERE outlet_id = $2
	`,
}
