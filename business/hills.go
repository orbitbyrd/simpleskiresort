package usecase

import (
	"skiresorts/models"
	"skiresorts/storage"
)

type IHillsService interface {
	GetHill(id int) (*models.Hill, error)
	UpdateHill(id int, hill *models.Hill) error
}

type hillsService struct {
	repo storage.IHillRepo
}

func NewHillService(repo storage.IHillRepo) *hillsService {
	return &hillsService{
		repo: repo,
	}
}

func (hs *hillsService) GetHill(id int) (*models.Hill, error) {
	return hs.repo.GetHill(id)
}

func (hs *hillsService) UpdateHill(id int, hill *models.Hill) error {
	return hs.repo.UpdateHill(id, hill)
}
