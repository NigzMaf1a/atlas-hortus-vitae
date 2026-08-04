package stock

type Stock struct {
	StockID   int64   `json:"stock_id"`
	OutletID  int64   `json:"outlet_id"`
	StockName string  `json:"stock_name"`
	StockQty  float64 `json:"stock_qty"`
	CostPerKg float64 `json:"cost_per_kg"`
}
