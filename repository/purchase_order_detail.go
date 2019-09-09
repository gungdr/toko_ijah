package repository

import (
	"toko_ijah/models"
)

//PurchaseOrderDetailRepository ..
type PurchaseOrderDetailRepository interface {
	Add(data *models.PurchaseOrderDetail) error
	GetByPurchaseOrderID(id int) []*models.PurchaseOrderDetail
}
