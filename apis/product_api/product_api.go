package product_api

import (
	"context"
	"encoding/json"
	"mockdb/config"
	"mockdb/entities"
	"mockdb/models"
	"mockdb/product"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	service interface {
		FindAll(ctx context.Context) (product []entities.Product, err error)
	}
	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		products, err2 := h.srv.FindAll(context.TODO())
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

// func FindAll(response http.ResponseWriter, request *http.Request) {
// 	db, err := config.GetDB()
// 	defer db.Close()
// 	if err != nil {
// 		respondWithError(response, http.StatusBadRequest, err.Error())
// 	} else {
// 		productModel := models.CreateProductModel(db)
// 		var service = product.NewService(productModel)
// 		products, err2 := service.FindAll(context.TODO())
// 		if err2 != nil {
// 			respondWithError(response, http.StatusBadRequest, err.Error())
// 		} else {
// 			respondWithJson(response, http.StatusOK, products)
// 		}
// 	}
// }

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.CreateProductModel(db)
		var service = product.NewService(productModel)
		products, err2 := service.Search(context.TODO(), keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var productitem entities.Product
	err := json.NewDecoder(request.Body).Decode(&productitem)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.CreateProductModel(db)
		var service = product.NewService(productModel)
		err2 := service.Create(context.TODO(), &productitem)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, productitem)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
