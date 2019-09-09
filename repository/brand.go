package repository

import (
	"toko_ijah/models"
)

//BrandRepository ...
type BrandRepository interface {
	Add(data *models.Brand) error
	GetByID(id int) *models.Brand
	GetByCode(code string) *models.Brand
}
