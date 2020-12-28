package main

import (
	"fmt"
	"mockdb/apis/product_api"
	"net/http"

	"github.com/gorilla/mux"
)

//index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home mock db!\n")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	//router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods("GET")
	router.HandleFunc("/api/product/create", product_api.Create).Methods("POST")
	fmt.Println("Listen port")
	err := http.ListenAndServe(":5000", router)
	if err != nil {

		fmt.Println(err)
	}
}
