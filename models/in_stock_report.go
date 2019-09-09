package models

import "time"

//InStockReport ..
type InStockReport struct {
	Time          time.Time `json:"time,omitempty" db:"time,omitempty"`
	SKU           string    `json:"sku,omitempty" db:"sku,omitempty"`
	ProductName   string    `json:"product_name,omitempty" db:"product_name,omitempty"`
	OrderQty      int       `json:"order_qty,omitempty" db:"order_qty,omitempty"`
	ReceiptQty    int       `json:"receipt_qty,omitempty" db:"receipt_qty,omitempty"`
	OrderPrice    float32   `json:"order_price,omitempty" db:"order_price,omitempty"`
	Total         float32   `json:"total,omitempty" db:"total,omitempty"`
	ReceiptNumber string    `json:"receipt_number,omitempty" db:"receipt_number,omitempty"`
	Notes         string    `json:"notes,omitempty" db:"notes,omitempty"`
}
