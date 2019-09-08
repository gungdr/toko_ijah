package models

import "time"

//ProductValueReport ..
type ProductValueReport struct {
	SKU         string  `json:"sku,omitempty" db:"sku,omitempty"`
	ProductName string  `json:"product_name,omitempty" db:"product_name,omitempty"`
	Qty         int     `json:"qty,omitempty" db:"qty,omitempty"`
	AvgPrice    float32 `json:"avg_price,omitempty" db:"avg_price,omitempty"`
	Total       float32 `json:"total,omitempty" db:"total,omitempty"`
}

//ProductValueReportSummary ..
type ProductValueReportSummary struct {
	PrintedDate time.Time             `json:"printed_date,omitempty"`
	SKUCount    int                   `json:"sku_count,omitempty"`
	TotalQty    int                   `json:"total_qty,omitempty"`
	TotalValue  float32               `json:"total_value,omitempty"`
	Detail      []*ProductValueReport `json:"detail,omitempty"`
}
