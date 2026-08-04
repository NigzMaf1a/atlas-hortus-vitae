package sales

import "time"

type Sale struct {
	SaleID       int64     `json:"sale_id"`
	UserId       int64     `json:"user_id"`
	OutletID     int64     `json:"outlet_id"`
	SaleDate     time.Time `json:"sale_date"`
	SaleTotal    float64   `json:"sale_total"`
	SaleDiscount float64   `json:"sale_discount"`
	SalePrice    float64   `json:"sale_price"`
}
