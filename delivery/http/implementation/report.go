package implementation

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"time"
	delivery "toko_ijah/delivery/http"
	"toko_ijah/infrastructure"
	tools "toko_ijah/infrastructure/tools"
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

func serveFormat(c *gin.Context, format string, report interface{}) {
	switch format {
	case infrastructure.FormatCSV:
		data := report.([][]string)
		buffer := &bytes.Buffer{}

		writer := csv.NewWriter(buffer)

		for _, value := range data {
			err := writer.Write(value)
			if err != nil {
				log.Println(err)
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		writer.Flush()

		c.Data(http.StatusOK, "text/csv", buffer.Bytes())
		return
	default:
		c.JSON(http.StatusOK, report)
		return
	}

}

//GetInStockReport ...
func (del *reportDelivery) GetInStockReport(c *gin.Context) {
	report := del.useCase.GetInStockReport()
	format := c.Query("format")

	if format == infrastructure.FormatCSV {
		lines := [][]string{{
			"Waktu",
			"SKU",
			"Nama Barang",
			"Jumlah Pemesanan",
			"Jumlah Diterima",
			"Harga Beli",
			"Total",
			"Nomer Kwitansi",
			"Catatan"}}
		for _, row := range report {
			lines = append(lines, []string{
				row.Time.Format(infrastructure.DateTimeFormat),
				row.SKU,
				row.ProductName,
				tools.ConvertNumberToString(row.OrderQty),
				tools.ConvertNumberToString(row.ReceiptQty),
				tools.ConvertNumberToString(row.OrderPrice),
				tools.ConvertNumberToString(row.Total),
				row.ReceiptNumber,
				row.Notes,
			},
			)
		}
		serveFormat(c, format, lines)
		return
	}

	serveFormat(c, infrastructure.FormatJSON, report)
}

//GetOutStockReport ...
func (del *reportDelivery) GetOutStockReport(c *gin.Context) {
	report := del.useCase.GetOutStockReport()
	format := c.Query("format")

	if format == infrastructure.FormatCSV {
		lines := [][]string{}
		for _, row := range report {
			lines = append(lines, []string{
				row.Time.Format(infrastructure.DateTimeFormat),
				row.SKU,
				row.ProductName,
				tools.ConvertNumberToString(row.Qty),
				tools.ConvertNumberToString(row.Price),
				tools.ConvertNumberToString(row.Total),
				row.Notes,
			},
			)
		}
		serveFormat(c, format, lines)
		return
	}
	serveFormat(c, infrastructure.FormatJSON, report)
}

func (del *reportDelivery) GetTotalProductReport(c *gin.Context) {
	report := del.useCase.GetTotalProductReport()
	format := c.Query("format")
	if format == infrastructure.FormatCSV {
		lines := [][]string{{"SKU", "Nama Item", "Jumlah Sekarang"}}
		for _, row := range report {
			lines = append(lines, []string{
				row.SKU,
				row.ProductName,
				tools.ConvertNumberToString(row.Qty),
			},
			)
		}
		serveFormat(c, format, lines)
		return
	}
	serveFormat(c, infrastructure.FormatJSON, report)
}

func (del *reportDelivery) GetProductValueReport(c *gin.Context) {
	report := del.useCase.GetProductValueReport()
	format := c.Query("format")
	if format == infrastructure.FormatCSV {
		lines := [][]string{{"LAPORAN NILAI BARANG"}}
		lines = append(lines, []string{})
		lines = append(lines, []string{"Tanggal Cetak", report.PrintedDate.Format(infrastructure.DateFormatFormal)})
		lines = append(lines, []string{"Jumlah SKU", tools.ConvertNumberToString(report.SKUCount)})
		lines = append(lines, []string{"Jumlah Total Barang", tools.ConvertNumberToString(report.TotalQty)})
		lines = append(lines, []string{"Total Nilai", tools.ConvertNumberToString(report.TotalValue)})
		lines = append(lines, []string{"SKU", "Nama Item", "Jumlah", "Rata-rata Harga Beli", "Total"})
		lines = append(lines, []string{})
		for _, row := range report.Detail {
			lines = append(lines, []string{
				row.SKU,
				row.ProductName,
				tools.ConvertNumberToString(row.Qty),
				tools.ConvertNumberToString(row.AvgPrice),
				tools.ConvertNumberToString(row.Total),
			},
			)
		}
		serveFormat(c, format, lines)
		return
	}

	serveFormat(c, infrastructure.FormatJSON, report)
}

func (del *reportDelivery) GetSalesReport(c *gin.Context) {

	yesterday := time.Now().Add(-24 * time.Hour).Format(infrastructure.DateFormat)
	today := time.Now().Format(infrastructure.DateFormat)
	fromParam := c.DefaultQuery("from", yesterday)
	untilParam := c.DefaultQuery("until", today)

	fromDate, err := time.Parse(infrastructure.DateFormat, fromParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	untilDate, err := time.Parse(infrastructure.DateFormat, untilParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	report := del.useCase.GetSalesReport(fromDate, untilDate)
	format := c.Query("format")
	if format == infrastructure.FormatCSV {
		lines := [][]string{{"LAPORAN PENJUALAN"}}
		lines = append(lines, []string{""})
		lines = append(lines, []string{"Tanggal Cetak", report.PrintedDate.Format(infrastructure.DateFormatFormal)})
		lines = append(lines, []string{"Tanggal",
			fmt.Sprintf("%s - %s",
				report.FromPeriod.Format(infrastructure.DateFormatFormal),
				report.UntilPeriod.Format(infrastructure.DateFormatFormal)),
		})
		lines = append(lines, []string{"Total Omzet", tools.ConvertNumberToString(report.Omzet)})
		lines = append(lines, []string{"Total Laba Kotor", tools.ConvertNumberToString(report.GrossProfit)})
		lines = append(lines, []string{"Total Penjualan", tools.ConvertNumberToString(report.TotalSales)})
		lines = append(lines, []string{"Total Barang", tools.ConvertNumberToString(report.TotalQty)})
		lines = append(lines, []string{""})
		lines = append(lines, []string{
			"ID Pesanan",
			"Waktu",
			"SKU",
			"Nama Barang",
			"Jumlah",
			"Harga Jual",
			"Total",
			"Harga Beli",
			"Laba",
		})
		for _, row := range report.Detail {
			lines = append(lines, []string{
				row.SalesID,
				row.Time.Format(infrastructure.DateTimeFormat),
				row.SKU,
				row.ProductName,
				tools.ConvertNumberToString(row.Qty),
				tools.ConvertNumberToString(row.SalesPrice),
				tools.ConvertNumberToString(row.SalesTotal),
				tools.ConvertNumberToString(row.OrderPrice),
				tools.ConvertNumberToString(row.Profit),
			},
			)
		}
		serveFormat(c, format, lines)
		return
	}
	serveFormat(c, infrastructure.FormatJSON, report)
}
