package usecase

import (
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
	report := useCase.repo.GetInStockReport()

	return report
}
