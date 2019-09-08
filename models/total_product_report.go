package models

//TotalProductReport ..
type TotalProductReport struct {
	SKU         string `json:"sku,omitempty" db:"sku,omitempty"`
	ProductName string `json:"product_name,omitempty" db:"product_name,omitempty"`
	Qty         int    `json:"qty,omitempty" db:"qty,omitempty"`
}
