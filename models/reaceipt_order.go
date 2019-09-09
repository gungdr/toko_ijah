package models

import "time"

//ReceiptOrder ..
type ReceiptOrder struct {
	ID            int       `json:"id,omitempty"`
	ReceiptNumber string    `json:"receipt_number,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Notes         string    `json:"notes,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	ModifiedOn    time.Time `json:"modified_on,omitempty"`
}
