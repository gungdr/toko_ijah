package models

import "time"

//Color ..
type Color struct {
	ID         int       `json:"id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}
