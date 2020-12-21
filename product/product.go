package product

import (
	"context"
	"mockdb/entities"
)

type (
	repository interface {
		FindAll(ctx context.Context) (product []entities.Product, err error)
	}
	Service struct {
		repo repository
	}
)

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindAll(ctx context.Context) (product []entities.Product, err error) {
	var products []entities.Product
	products, err = s.repo.FindAll(ctx)
	if err != nil {
		return products, err
	}
	return products, nil
}
