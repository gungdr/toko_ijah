package implementation

import (
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

INNER JOIN
	(
	SELECT 
	product.id,
	'` + repo.prefix + `-' || brand.code || product.id || '-' || "size".code || '-' || color.code as sku,
	brand.name ||' '|| product.name ||' ('|| "size".name ||', '|| color.name || ')' as name
	from product
	INNER JOIN "size"
	on product.size_id = "size".id
	INNER join brand
	on product.brand_id = brand.id
	INNER join color
	on product.color_id = color.id) prod
on prod.id = ro.product_id
and prod.id = po.product_id`
	var report []*models.InStockReport
	err := repo.db.Select(&report, query)
	return report, err
}
