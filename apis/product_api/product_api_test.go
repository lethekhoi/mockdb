package product_api_test

import (
	"mockdb/apis/product_api"
	"mockdb/entities"
	"mockdb/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandlerFindAll(t *testing.T) {
	noproduct := []entities.Product{}
	products := []entities.Product{
		{
			Id:       1,
			Name:     "1",
			Price:    1,
			Quantity: 1,
		},
	}
	ctr := gomock.NewController(t)
	mockService := product_api.NewMockservice(ctr)
	handler := product_api.NewHandler(mockService)
	type expect struct {
		code int
	}
	testCases := []struct {
		name     string
		tearDown func()
		expect   expect
	}{
		{
			name: "find success",
			tearDown: func() {
				mockService.EXPECT().FindAll(gomock.Any()).Return(products, nil).Times(1)
			},

			expect: expect{
				code: http.StatusOK,
			},
		},
		{
			name: "find fail",
			tearDown: func() {
				mockService.EXPECT().FindAll(gomock.Any()).Return(noproduct, errors.ErrNotFound).Times(1)
			},

			expect: expect{
				code: http.StatusBadRequest,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Error(err)
			}
			handler.FindAll(w, r)
			if w.Code != test.expect.code {
				t.Errorf("got code=%d, wants code=%d", w.Code, test.expect.code)
			}
		})
	}

}
