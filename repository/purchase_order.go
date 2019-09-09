package repository

import (
	"time"
	"toko_ijah/models"
)

//PurchaseOrderRepository ..
type PurchaseOrderRepository interface {
	Add(data *models.PurchaseOrder) error
	GetByID(id int) *models.PurchaseOrder
	GetByDate(from, until time.Time) *models.PurchaseOrder
}
