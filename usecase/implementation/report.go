package usecase

import (
	"log"
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
