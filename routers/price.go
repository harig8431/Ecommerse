package routers

import (
	"net/http"
	"project1/Controllers/masters"

	"github.com/gorilla/mux"
)

func PriceRoutes(router *mux.Router) *mux.Router {
	pricecontroller := masters.Pricecontroller{}
	router.Handle("/price/create", http.HandlerFunc(pricecontroller.CreatePrice)).Methods(http.MethodPost)
	router.Handle("/price/update", http.HandlerFunc(pricecontroller.UpdatePrice)).Methods(http.MethodPost)
	router.Handle("/price/getbyid/{id}", http.HandlerFunc(pricecontroller.GetPricebyid)).Methods(http.MethodGet)
	router.Handle("/price/getall", http.HandlerFunc(pricecontroller.GetAllPrices)).Methods(http.MethodGet)

	return router
}
