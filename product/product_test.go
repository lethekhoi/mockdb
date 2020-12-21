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

	products := []entities.Product{
		{
			Id:       1,
			Name:     "1",
			Price:    1,
			Quantity: 1,
		},
	}
	service := product.NewService(repoMocked)
	repoMocked.EXPECT().FindAll(gomock.Any()).Return(products, nil).Times(1)
	results, err := service.FindAll(context.TODO())
	if err != nil {
		t.Errorf("got err = %v, expects err = %v", err, nil)
	}
	if len(results) != len(products) {
		t.Errorf("got len = %v, expects len = %v", len(products), len(products))
	}

}
