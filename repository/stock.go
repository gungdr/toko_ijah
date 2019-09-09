package repository

import "toko_ijah/models"

//StockRepository ..
type StockRepository interface {
	Add(data *models.Stock) error
	GetByID(id int) *models.Stock
	GetByReceiptOrder(id int) []*models.Stock
	GetByProductID(id int) []*models.Stock
}
