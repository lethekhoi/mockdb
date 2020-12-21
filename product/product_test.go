package product_test

import (
	"context"
	"mockdb/entities"
	"mockdb/product"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFindAll(t *testing.T) {
	ctr := gomock.NewController(t)
	repoMocked := product.NewMockrepository(ctr)
	service := product.NewService(repoMocked)
	products := []entities.Product{
		{
			Id:       1,
			Name:     "1",
			Price:    1,
			Quantity: 1,
		},
	}

	testCases := []struct {
		name     string
		products []entities.Product
		tearDown func()
		ouput    error
	}{
		{
			name:     "Find All Product",
			products: products,
			tearDown: func() {
				repoMocked.EXPECT().FindAll(gomock.Any()).Times(1).Return(products, nil)

			},
			ouput: nil,
		},
		{
			name:     "Not Found Any Product",
			products: []entities.Product{},
			tearDown: func() {
				repoMocked.EXPECT().FindAll(gomock.Any()).Times(1).Return([]entities.Product{}, product.ErrNotFound)

			},
			ouput: product.ErrNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			products, err := service.FindAll(context.TODO())
			if err != test.ouput {
				t.Errorf("got err = %v, expects err = %v", err, test.ouput)
			}
			if len(products) != len(test.products) {
				t.Errorf("got len = %v, expects len = %v", len(products), len(test.products))
			}
		})
	}

}
func TestSearch(t *testing.T) {
	ctr := gomock.NewController(t)
	repoMocked := product.NewMockrepository(ctr)
	service := product.NewService(repoMocked)
	products := []entities.Product{
		{
			Id:       1,
			Name:     "1",
			Price:    1,
			Quantity: 1,
		},
	}
	testCases := []struct {
		name     string
		keyword  string
		products []entities.Product
		tearDown func()
		ouput    error
	}{
		{
			name:     "Find By keyword",
			products: products,
			keyword:  "1",
			tearDown: func() {

				repoMocked.EXPECT().Search(gomock.Any(), "1").Times(1).Return(products, nil)

			},
			ouput: nil,
		},
		{
			name:     "Find By keyword no result",
			keyword:  "1",
			products: []entities.Product{},
			tearDown: func() {

				repoMocked.EXPECT().Search(gomock.Any(), "1").Times(1).Return([]entities.Product{}, product.ErrNotFound)

			},
			ouput: product.ErrNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			products, err := service.Search(context.TODO(), test.keyword)
			if err != test.ouput {
				t.Errorf("got err = %v, expects err = %v", err, test.ouput)
			}
			if len(products) != len(test.products) {
				t.Errorf("got len = %v, expects len = %v", len(products), len(test.products))
			}
		})
	}
}
