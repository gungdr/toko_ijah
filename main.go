package main

import (
	"fmt"
	"log"
	"os"
	delivery "toko_ijah/delivery/http/implementation"
	"toko_ijah/infrastructure"
	repository "toko_ijah/repository/implementation/sqlite"
	usecase "toko_ijah/usecase/implementation"

	"github.com/gin-gonic/gin"
)

func main() {
	prefix := os.Getenv("SKU_PREFIX")
	driver := os.Getenv("DB_DRIVER")
	connString := os.Getenv("DB_CONNECTION")
	htptPort := os.Getenv("HTTP_PORT")
	db := infrastructure.InitDB(driver, connString)
	defer db.Close()

	log.Println(prefix)
	reportRepo := repository.NewReportRepository(db, prefix)
	reportUseCase := usecase.NewReportUseCase(reportRepo)
	reportDelivery := delivery.NewReportDelivery(reportUseCase)

	r := gin.Default()
	reportRoutes := r.Group("/report")
	{
		reportRoutes.GET("/in-stock", reportDelivery.GetInStockReport)
		reportRoutes.GET("/out-stock", reportDelivery.GetOutStockReport)
		reportRoutes.GET("/total-product", reportDelivery.GetTotalProductReport)
		reportRoutes.GET("/product-value", reportDelivery.GetProductValueReport)
	}
	r.Run(fmt.Sprintf(":%s", htptPort))

}
