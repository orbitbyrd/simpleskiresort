package persistence

import (
	"fmt"
	"skiresorts/domain"
)

type inMemDB struct {
	Hills map[int]*domain.Hill
}

func newInMemDB() *inMemDB {
	hills := map[int]*domain.Hill{
		1: {
			Length: 91.1,
			Slope:  0.29,
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

func (imr *inMemRepo) GetHill(id int) (*domain.Hill, error) {
	if h, ok := imr.db.Hills[id]; ok {
		return h, nil
	}
	return nil, fmt.Errorf("Hill with id %d is not found", id)
}

func (imr *inMemRepo) UpdateHill(id int, hill *domain.Hill) error {
	if _, ok := imr.db.Hills[id]; ok {
		imr.db.Hills[id] = hill
		return nil
	}
	return fmt.Errorf("Hill with id %d not found", id)
}
