package repository

import (
	"time"
	"toko_ijah/models"
)

//ReceiptOrderRepository ..
type ReceiptOrderRepository interface {
	Add(data *models.ReceiptOrder) error
	GetByID(id int) *models.ReceiptOrder
	GetByDate(from, until time.Time) *models.ReceiptOrder
	GetByPurchaseOrderID(id int) []*models.ReceiptOrder
}
