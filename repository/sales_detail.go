package repository

import (
	"toko_ijah/models"
)

//SalesDetailRepository ..
type SalesDetailRepository interface {
	Add(data *models.SalesDetail) error
	GetBySalesID(id int) []*models.SalesDetail
}
