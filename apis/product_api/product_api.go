package product_api

import (
	"context"
	"encoding/json"
	"mockdb/config"
	"mockdb/models"
	"mockdb/product"
	"net/http"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.CreateProductModel(db)
		var service = product.NewService(productModel)
		products, err2 := service.FindAll(context.TODO())
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(response, http.StatusOK, products)
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
