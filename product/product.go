package product

import (
	"context"
	"fmt"
	"mockdb/entities"
)

type (
	repository interface {
		FindAll(ctx context.Context) (product []entities.Product, err error)
		Search(ctx context.Context, ID string) (product []entities.Product, err error)
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

func (s *Service) Search(ctx context.Context, keyword string) (products []entities.Product, err error) {
	fmt.Println("search")
	products, err = s.repo.Search(ctx, keyword)
	if err != nil {
		return products, err
	}
	return products, nil
}
