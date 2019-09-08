package usecase

import (
	"log"
	"time"
	"toko_ijah/models"
	"toko_ijah/repository"
	"toko_ijah/usecase"
)

type reportUseCase struct {
	repo repository.ReportRepository
}

//NewReportUseCase ..
func NewReportUseCase(repo repository.ReportRepository) usecase.ReportUseCase {
	return &reportUseCase{
		repo,
	}
}

func (useCase *reportUseCase) GetInStockReport() []*models.InStockReport {
	report, err := useCase.repo.GetInStockReport()
	if err != nil {
		log.Println(err)
	}
	return report
}

func (useCase *reportUseCase) GetOutStockReport() []*models.OutStockReport {
	report, err := useCase.repo.GetOutStockReport()
	if err != nil {
		log.Println(err)
	}
	return report
}

func (useCase *reportUseCase) GetTotalProductReport() []*models.TotalProductReport {
	report, err := useCase.repo.GetTotalProductReport()
	if err != nil {
		log.Println(err)
	}
	return report
}

func (useCase *reportUseCase) GetProductValueReport() *models.ProductValueReportSummary {
	report, err := useCase.repo.GetProductValueReport()
	type summaryMap struct {
		qty   int
		value float32
	}
	skuCount := map[string]bool{}
	qty := 0
	value := float32(0)
	for _, row := range report {
		qty = qty + row.Qty
		value = value + row.Total
		skuCount[row.SKU] = true
	}
	summary := &models.ProductValueReportSummary{
		PrintedDate: time.Now(),
		SKUCount:    len(skuCount),
		TotalQty:    qty,
		TotalValue:  value,
		Detail:      report,
	}
	if err != nil {
		log.Println(err)
	}
	return summary
}
