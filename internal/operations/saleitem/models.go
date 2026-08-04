package saleitem

type SaleItem struct {
	SaleItemID   int64 `json:"sale_item_id"`
	SaleID       int64 `json:"sale_id"`
	ProductID    int64 `json:"product_id"`
	ProductPrice int64 `json:"product_price"`
	SaleQty      int64 `json:"sale_qty"`
	SaleTotal    int64 `json:"sale_total"`
}
