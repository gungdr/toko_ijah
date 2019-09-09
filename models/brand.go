package models

import "time"

//Brand ..
type Brand struct {
	ID        int       `json:"id,omitempty"`
	Code      string    `json:"code,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
}
