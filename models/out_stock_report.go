package models

import "time"

//OutStockReport ..
type OutStockReport struct {
	Time        time.Time `json:"time,omitempty" db:"time,omitempty"`
	SKU         string    `json:"sku,omitempty" db:"sku,omitempty"`
	ProductName string    `json:"product_name,omitempty" db:"product_name,omitempty"`
	Qty         int       `json:"qty,omitempty" db:"qty,omitempty"`
	Price       float32   `json:"price,omitempty" db:"price,omitempty"`
	Total       float32   `json:"total,omitempty" db:"total,omitempty"`
	Notes       string    `json:"notes,omitempty" db:"notes,omitempty"`
}
