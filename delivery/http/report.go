package delivery

import "github.com/gin-gonic/gin"

//ReportDelivery ..
type ReportDelivery interface {
	GetInStockReport(c *gin.Context)
	GetOutStockReport(c *gin.Context)
}
