package models

import "time"

//ReceiptOrderDetail ..
type ReceiptOrderDetail struct {
	ID           int          `json:"id,omitempty"`
	ReceiptOrder ReceiptOrder `json:"receipt_order,omitempty"`
	Product      Product      `json:"product,omitempty"`
	CreatedOn    time.Time    `json:"created_on,omitempty"`
	ModifiedOn   time.Time    `json:"modified_on,omitempty"`
}
