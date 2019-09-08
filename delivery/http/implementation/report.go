package implementation

import (
	"net/http"
	delivery "toko_ijah/delivery/http"
	"toko_ijah/usecase"

	"github.com/gin-gonic/gin"
)

type reportDelivery struct {
	useCase usecase.ReportUseCase
}

//NewReportDelivery ..
func NewReportDelivery(useCase usecase.ReportUseCase) delivery.ReportDelivery {
	return &reportDelivery{
		useCase,
	}
}

//GetInStockReport ...
func (del *reportDelivery) GetInStockReport(c *gin.Context) {
	report := del.useCase.GetInStockReport()
	c.JSON(http.StatusOK, report)
}
