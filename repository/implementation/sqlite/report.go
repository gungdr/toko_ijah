package implementation

import (
	"time"
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

func (repo *reportRepository) GetSalesReport(from, until time.Time) ([]*models.SalesReport, error) {
	query := `
			select 
			'ID-'|| STRFTIME('%Y%m%d',sales.date) ||'-'||  sales.id as sales_id ,
			sales.date as "time",
			vw_product.sku,
			vw_product.name as product_name,
			sales_detail.price as sales_price,
			vw_in_stock_report.order_price / vw_in_stock_report.order_qty as order_price,
			sales_sub.qty,
			sales_sub.qty * sales_detail.price as sales_total,
			(sales_sub.qty * sales_detail.price) - (vw_in_stock_report.order_price / vw_in_stock_report.order_qty * sales_sub.qty) as profit 
			FROM sales_detail
			INNER join sales
			on sales.id = sales_detail.sales_id
			INNER join stock
			on stock.id = sales_detail.stock_id
			INNER join vw_product
			on stock.product_id = vw_product.id
			INNER join (
				SELECT 
				sales_id,
				count(stock_id) as qty
				from sales_detail
				group by sales_id
			) sales_sub
			
			on sales_sub.sales_id = sales.id
			
			left JOIN vw_in_stock_report
			on vw_in_stock_report.sku = vw_product.sku
			
			WHERE sales.date BETWEEN $1 and $2
			
			group by
			sales.id,
			sales.date,
			vw_product.sku,
			vw_product.name
			order by sales.id	
	`
	var report []*models.SalesReport
	err := repo.db.Select(&report, query, from, until)

	return report, err

}

func (repo *reportRepository) GetProductValueReport() ([]*models.ProductValueReport, error) {
	query := `
	
		select
		vw_product.sku,
		vw_product.name as product_name,
		count(stock.product_id) as qty,
		po.price as avg_price,
		count(stock.product_id) * po.price as total
		from stock
		inner join vw_product
		on vw_product.id = stock.product_id
		inner join (
			select 
			purchase_order_detail.product_id,
			avg(purchase_order_detail.price) as price
			from purchase_order_detail
			group by
			purchase_order_detail.product_id
		)po
		on po.product_id = stock.product_id

		WHERE stock.id not in (
			SELECT 
			stock.id
			from stock
			inner join sales_detail
			on sales_detail.stock_id = stock.id
		)

		GROUP by stock.product_id
	`

	var report []*models.ProductValueReport
	err := repo.db.Select(&report, query)
	return report, err
}

func (repo *reportRepository) GetTotalProductReport() ([]*models.TotalProductReport, error) {
	query := `
			select
			vw_product.sku,
			vw_product.name as product_name,
			count(stock.product_id) as qty
			
			from stock
			inner join vw_product
			on vw_product.id = stock.product_id
			
			WHERE stock.id not in (
				SELECT 
				stock.id
				from stock
				inner join sales_detail
				on sales_detail.stock_id = stock.id
			)

			GROUP by stock.product_id
	`
	var report []*models.TotalProductReport
	err := repo.db.Select(&report, query)
	return report, err
}

func (repo *reportRepository) GetOutStockReport() ([]*models.OutStockReport, error) {
	query := `
		select 
		sales.date as "time",
		prod.sku,
		prod.name as product_name,
		count(stock.product_id) as qty,
		avg(sales_detail.price) over(PARTITION by sales.id,sales.date) as price,
		sum(sales_detail.price) as total,
		'ID-'|| STRFTIME('%Y%m%d',sales.date) ||'-'||  sales.id as notes
		from sales_detail
		
		inner JOIN sales
		on sales.id = sales_detail.sales_id
		inner join stock
		on stock.id = sales_detail.stock_id

		INNER join vw_product prod prod
		on prod.id = stock.product_id
		group BY
		sales.id,
		sales.date
	`
	var report []*models.OutStockReport
	err := repo.db.Select(&report, query)
	return report, err
}

func (repo *reportRepository) GetInStockReport() ([]*models.InStockReport, error) {
	query := `
				select
				ro.date as "time",
				prod.sku as sku,
				prod.name as product_name,
				po.order_qty as order_qty,
				ro.qty as receipt_qty,
				po.sum_price as order_price,
				coalesce(ro.receipt_number, '') as receipt_number,
				coalesce(STRFTIME('%Y/%m/%d',ro.date) ||' terima ' || ro.qty, '' )as notes
			from
				(
				select
					purchase_order.id,
					purchase_order.date,
					purchase_order_detail.product_id,
					count(*) order_qty,
					purchase_order_detail.price,
					sum(purchase_order_detail.price) as sum_price
				from
					purchase_order
				INNER join purchase_order_detail on
					purchase_order.id = purchase_order_detail.purchase_order_id
				group by
					purchase_order.date,
					purchase_order_detail.product_id,
					purchase_order_detail.price ) po
			inner join (
				SELECT
					receipt_order.purchase_order_id,
					receipt_order.receipt_number,
					receipt_order.date,
					receipt_order_detail.product_id,
					count(*) qty
				from
					receipt_order
				INNER join receipt_order_detail on
					receipt_order.id = receipt_order_detail.receipt_order_id
				group by
					receipt_order.receipt_number,
					receipt_order.date,
					receipt_order_detail.product_id,
					receipt_order.purchase_order_id )ro on
				ro.purchase_order_id = po.id
				and ro.product_id = po.product_id

			INNER JOIN vw_product prod
			on prod.id = ro.product_id
			and prod.id = po.product_id
	`
	var report []*models.InStockReport
	err := repo.db.Select(&report, query)
	return report, err
}
