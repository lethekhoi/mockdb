package product

import (
	"context"
	"mockdb/entities"
)

type (
	repository interface {
		FindAll(ctx context.Context) (product []entities.Product, err error)
		//	SearchID(ctx context.Context, ID string) (product entities.Product, err error)
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

// func (s *Service) SearchID(ctx context.Context, ID string) (product entities.Product, err error) {

// 	product, err = s.repo.SearchID(ctx, ID)
// 	if err != nil {
// 		return product, err
// 	}
// 	return product, nil
// }
