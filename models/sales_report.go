package models

import "time"

//SalesReport ..
type SalesReport struct {
	SalesID     string    `json:"sales_id,omitempty" db:"sales_id,omitempty"`
	Time        time.Time `json:"time,omitempty" db:"time,omitempty"`
	SKU         string    `json:"sku,omitempty" db:"sku,omitempty"`
	ProductName string    `json:"product_name,omitempty" db:"product_name,omitempty"`
	Qty         int       `json:"qty,omitempty" db:"qty,omitempty"`
	SalesPrice  float32   `json:"sales_price,omitempty" db:"sales_price,omitempty"`
	SalesTotal  float32   `json:"sales_total,omitempty" db:"sales_total,omitempty"`
	OrderPrice  float32   `json:"order_price,omitempty" db:"order_price,omitempty"`
	Profit      float32   `json:"profit,omitempty" db:"profit,omitempty"`
}

//SalesReportSummary ..
type SalesReportSummary struct {
	PrintedDate time.Time      `json:"printed_date,omitempty"`
	FromPeriod  time.Time      `json:"from_period,omitempty"`
	UntilPeriod time.Time      `json:"until_period,omitempty"`
	Omzet       float32        `json:"omzet,omitempty"`
	GrossProfit float32        `json:"gross_profit,omitempty"`
	TotalSales  int            `json:"total_sales,omitempty"`
	TotalQty    int            `json:"total_qty,omitempty"`
	Detail      []*SalesReport `json:"detail,omitempty"`
}
