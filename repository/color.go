package repository

import (
	"toko_ijah/models"
)

//ColorRepository ...
type ColorRepository interface {
	Add(data *models.Color) error
	GetByID(id int) *models.Color
	GetByCode(code string) *models.Color
}
