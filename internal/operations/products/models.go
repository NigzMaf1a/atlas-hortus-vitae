package products

type Product struct {
	ProductID    int64  `json:"product_id"`
	OutletID     int64  `json:"outlet_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int64  `json:"product_price"`
	Available    string `json:"available"`
}
