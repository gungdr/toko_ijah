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
	if err != nil {
		log.Println(err)
		return nil
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

	return summary
}

func (useCase *reportUseCase) GetSalesReport(from, until time.Time) *models.SalesReportSummary {
	report, err := useCase.repo.GetSalesReport(from, until)
	if err != nil {
		log.Println(err)
		return nil
	}
	totalOmzet, totalGross := float32(0), float32(0)
	totalQty := 0
	salesCount := map[string]bool{}
	for _, row := range report {
		totalOmzet = totalOmzet + row.SalesTotal
		totalGross = totalGross + row.Profit
		totalQty = totalQty + row.Qty
		salesCount[row.SalesID] = true
	}

	summary := &models.SalesReportSummary{
		PrintedDate: time.Now(),
		FromPeriod:  from,
		UntilPeriod: until,
		Omzet:       totalOmzet,
		GrossProfit: totalGross,
		TotalSales:  len(salesCount),
		TotalQty:    totalQty,
		Detail:      report,
	}

	return summary
}
