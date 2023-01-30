package routers

import (
	"net/http"
	
	"project/Controllers/masters"

	"github.com/gorilla/mux"
)

func CategoryRoutes(router *mux.Router) *mux.Router {
	categorycontroller := masters.Categorycontroller{}
	router.Handle("/category/create", http.HandlerFunc(categorycontroller.CreateCategory)).Methods(http.MethodPost)
	router.Handle("/category/update", http.HandlerFunc(categorycontroller.UpdateCategory)).Methods(http.MethodPost)
	router.Handle("/category/getbyid/{id}", http.HandlerFunc(categorycontroller.GetCategorybyid)).Methods(http.MethodGet)
	router.Handle("/category/getall", http.HandlerFunc(categorycontroller.GetAllCategorys)).Methods(http.MethodGet)

	return router
}
