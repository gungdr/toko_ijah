package implementation

import (
	"log"
	"toko_ijah/models"
	"toko_ijah/repository"

	"github.com/jmoiron/sqlx"
)

type reportRepository struct {
	prefix string
	db     *sqlx.DB
}

//NewReportRepository ...
func NewReportRepository(db *sqlx.DB, prefix string) repository.ReportRepository {
	return &reportRepository{
		prefix,
		db,
	}
}

func (repo *reportRepository) GetInStockReport() []*models.InStockReport {
	query := ``
	var report []*models.InStockReport
	err := repo.db.Select(&report, query)
	if err != nil {
		log.Println(err)
		return nil
	}
	return report
}
