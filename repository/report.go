package repository

import (
	"time"
	"toko_ijah/models"
)

//ReportRepository ..
type ReportRepository interface {
	GetInStockReport() ([]*models.InStockReport, error)
	GetOutStockReport() ([]*models.OutStockReport, error)
	GetTotalProductReport() ([]*models.TotalProductReport, error)
	GetProductValueReport() ([]*models.ProductValueReport, error)
	GetSalesReport(from, until time.Time) ([]*models.SalesReport, error)
}
