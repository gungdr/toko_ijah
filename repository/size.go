package repository

import "toko_ijah/models"

//SizeRepository ..
type SizeRepository interface {
	Add(data *models.Size) error
	GetSizeByID(id int) *models.Size
	GetSizeByCode(code string) *models.Size
}
