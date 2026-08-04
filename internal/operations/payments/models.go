package payments

import "time"

type Payment struct {
	PaymentID     int64     `json:"payment_id"`
	UserId        int64     `json:"user_id"`
	PaymentDate   time.Time `json:"payment_date"`
	SaleID        int64     `json:"sale_id"`
	SalePrice     float64   `json:"sale_price"`
	PaymentCode   string    `json:"payment_code"`
	PaymentStatus string    `json:"payment_status"`
}
