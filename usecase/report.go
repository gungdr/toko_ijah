package usecase

import (
	"time"
	"toko_ijah/models"
)

//ReportUseCase ..
type ReportUseCase interface {
	GetInStockReport() []*models.InStockReport
	GetOutStockReport() []*models.OutStockReport
	GetTotalProductReport() []*models.TotalProductReport
	GetProductValueReport() *models.ProductValueReportSummary
	GetSalesReport(from, until time.Time) *models.SalesReportSummary
}
