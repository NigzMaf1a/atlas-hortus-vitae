package outlets

type Outlet struct {
	OutletID int64   `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Networth float64 `json:"networth"`
	Open     string  `json:"open"`
}
