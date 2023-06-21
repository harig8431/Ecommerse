package routers

import (
	"net/http"
	"project1/Controllers"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) *mux.Router {
	productcontroller := Controllers.Productcontroller{}
	router.Handle("/product/create", http.HandlerFunc(productcontroller.CreateProduct)).Methods(http.MethodPost)
	router.Handle("/product/update", http.HandlerFunc(productcontroller.UpdateProduct)).Methods(http.MethodPost)
	router.Handle("/product/getbyid/{id}", http.HandlerFunc(productcontroller.GetProductbyid)).Methods(http.MethodGet)
	router.Handle("/product/getall", http.HandlerFunc(productcontroller.GetAllProducts)).Methods(http.MethodGet)
	router.Handle("/product/search/{key}", http.HandlerFunc(productcontroller.ProductSearch)).Methods(http.MethodGet)

	return router
}
