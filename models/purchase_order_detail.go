package models

import "time"

//PurchaseOrderDetail ..
type PurchaseOrderDetail struct {
	ID            int           `json:"id,omitempty"`
	PurchaseOrder PurchaseOrder `json:"purchase_order,omitempty"`
	Product       Product       `json:"product,omitempty"`
	Price         float32       `json:"price,omitempty"`
	CreatedOn     time.Time     `json:"created_on,omitempty"`
	ModifiedOn    time.Time     `json:"modified_on,omitempty"`
}
