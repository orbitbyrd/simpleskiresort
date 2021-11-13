package business

import (
	"skiresorts/domain"
	"skiresorts/persistence"
)

type IHillsService interface {
	GetHill(id int) (*domain.Hill, error)
	UpdateHill(id int, hill *domain.Hill) error
}

type hillsService struct {
	repo persistence.IHillRepo
}

func NewHillService(repo persistence.IHillRepo) *hillsService {
	return &hillsService{
		repo: repo,
	}
}

func (hs *hillsService) GetHill(id int) (*domain.Hill, error) {
	return hs.repo.GetHill(id)
}

func (hs *hillsService) UpdateHill(id int, hill *domain.Hill) error {
	return hs.repo.UpdateHill(id, hill)
}
