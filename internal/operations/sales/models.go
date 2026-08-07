package sales

import (
	"time"

	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/operations/saleitem"
)

type Sale struct {
	SaleID       int64     `json:"sale_id"`
	UserId       int64     `json:"user_id"`
	OutletID     int64     `json:"outlet_id"`
	SaleDate     time.Time `json:"sale_date"`
	SaleStatus   string    `json:"sale_status"`
	SaleTotal    float64   `json:"sale_total"`
	SaleDiscount float64   `json:"sale_discount"`
	SalePrice    float64   `json:"sale_price"`
}

type CreateSaleRequest struct {
	Sale  Sale                `json:"sale"`
	Items []saleitem.SaleItem `json:"items"`
}
