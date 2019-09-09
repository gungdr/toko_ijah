package repository

import "toko_ijah/models"

//ProductRepository ..
type ProductRepository interface {
	Add(data *models.Product) error
	GetProductByID(id int) *models.Product
	GetProductByCode(code string) *models.Product
}
