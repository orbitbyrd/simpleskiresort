package persistence

import "skiresorts/domain"

type IHillRepo interface {
	GetHill(id int) (*domain.Hill, error)
	UpdateHill(id int, hill *domain.Hill) error
}
