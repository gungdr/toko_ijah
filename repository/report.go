package repository

import "toko_ijah/models"

//ReportRepository ..
type ReportRepository interface {
	GetInStockReport() []*models.InStockReport
}