package models

import "time"

//Sales ..
type Sales struct {
	ID         int       `json:"id,omitempty"`
	Date       time.Time `json:"date,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	IsVoid     bool      `json:"is_void,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}
