package models

import "time"

//Stock ..
type Stock struct {
	ID           int          `json:"id,omitempty"`
	ReceiptOrder ReceiptOrder `json:"receipt_order,omitempty"`
	Product      Product      `json:"product,omitempty"`
	Price        float32      `json:"price,omitempty"`
	CreatedOn    time.Time    `json:"created_on,omitempty"`
	ModifiedOn   time.Time    `json:"modified_on,omitempty"`
}
