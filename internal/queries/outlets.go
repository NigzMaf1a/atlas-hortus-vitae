package queries

type OutletQuery struct {
	Create             string
	ReadOutlets        string
	ReadOutlet         string
	UpdateNetworth     string
	UpdateOutletStatus string
}

var OutletQueries OutletQuery = OutletQuery{
	Create:             ``,
	ReadOutlets:        ``,
	ReadOutlet:         ``,
	UpdateNetworth:     ``,
	UpdateOutletStatus: ``,
}
