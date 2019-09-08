package implementation

import (
	"net/http"
	"time"
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

//GetOutStockReport ...
func (del *reportDelivery) GetOutStockReport(c *gin.Context) {
	report := del.useCase.GetOutStockReport()
	c.JSON(http.StatusOK, report)
}

func (del *reportDelivery) GetTotalProductReport(c *gin.Context) {
	report := del.useCase.GetTotalProductReport()
	c.JSON(http.StatusOK, report)
}

func (del *reportDelivery) GetProductValueReport(c *gin.Context) {
	report := del.useCase.GetProductValueReport()
	c.JSON(http.StatusOK, report)
}

func (del *reportDelivery) GetSalesReport(c *gin.Context) {
	format := "2006-01-02"
	yesterday := time.Now().Add(-24 * time.Hour).Format(format)
	today := time.Now().Format(format)
	fromParam := c.DefaultQuery("from", yesterday)
	untilParam := c.DefaultQuery("until", today)

	fromDate, err := time.Parse(format, fromParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	untilDate, err := time.Parse(format, untilParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	report := del.useCase.GetSalesReport(fromDate, untilDate)
	c.JSON(http.StatusOK, report)
}
