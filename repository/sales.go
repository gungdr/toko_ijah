package repository

import (
	"time"
	"toko_ijah/models"
)

//SalesRepository ..
type SalesRepository interface {
	Add(data *models.Sales) error
	GetByID(id int) *models.Sales
	GetByDate(from, until time.Time) []*models.Sales
}
