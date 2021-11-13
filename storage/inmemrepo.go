package storage

import (
	"fmt"
	"skiresorts/models"
)

type inMemDB struct {
	Hills map[int]*models.Hill
}

func newInMemDB() *inMemDB {
	hills := map[int]*models.Hill{
		1: {
			Length: 99.9,
			Slope:  0.21,
		},
	}
	return &inMemDB{
		Hills: hills,
	}
}

type inMemRepo struct {
	db *inMemDB
}

func NewInMemRepo() *inMemRepo {
	return &inMemRepo{
		db: newInMemDB(),
	}
}

func (imr *inMemRepo) GetHill(id int) (*models.Hill, error) {
	if h, ok := imr.db.Hills[id]; ok {
		return h, nil
	}
	return nil, fmt.Errorf("Hill with id %d is not found", id)
}

func (imr *inMemRepo) UpdateHill(id int, hill *models.Hill) error {
	if _, ok := imr.db.Hills[id]; ok {
		imr.db.Hills[id] = hill
		return nil
	}
	return fmt.Errorf("Hill with id %d not found", id)
}
