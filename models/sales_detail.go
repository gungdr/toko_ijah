package models

import "time"

//SalesDetail ..
type SalesDetail struct {
	ID         int       `json:"id,omitempty"`
	Sales      Sales     `json:"sales,omitempty"`
	Stock      Stock     `json:"stock,omitempty"`
	Price      float32   `json:"price,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}
