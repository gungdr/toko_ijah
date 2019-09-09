package models

import "time"

//Product ..
type Product struct {
	ID         int       `json:"id,omitempty"`
	Brand      Brand     `json:"brand,omitempty"`
	Color      Color     `json:"color,omitempty"`
	Size       Size      `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}
