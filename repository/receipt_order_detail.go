package repository

import (
	"toko_ijah/models"
)

//ReceiptOrderDetailRepository ..
type ReceiptOrderDetailRepository interface {
	Add(data *models.ReceiptOrderDetail) error
	GetByReceiptID(id int) []*models.ReceiptOrderDetail
}
