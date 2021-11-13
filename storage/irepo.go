package storage

import "skiresorts/models"

type IHillRepo interface {
	GetHill(id int) (*models.Hill, error)
	UpdateHill(id int, hill *models.Hill) error
}
