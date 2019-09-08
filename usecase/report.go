package usecase

import "toko_ijah/models"

//ReportUseCase ..
type ReportUseCase interface {
	GetInStockReport() []*models.InStockReport
	GetOutStockReport() []*models.OutStockReport
	GetTotalProductReport() []*models.TotalProductReport
	GetProductValueReport() *models.ProductValueReportSummary
}
