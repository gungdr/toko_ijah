package models

import "time"

//PurchaseOrder ..
type PurchaseOrder struct {
	ID         int       `json:"id,omitempty"`
	Date       time.Time `json:"date,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}
