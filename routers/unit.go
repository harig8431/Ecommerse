package routers

import (
	"net/http"

	"project/Controllers/masters"

	"github.com/gorilla/mux"
)

func UnitRoutes(router *mux.Router) *mux.Router {
	unitcontroller := masters.Unitcontroller{}
	router.Handle("/unit/create", http.HandlerFunc(unitcontroller.CreateUnit)).Methods(http.MethodPost)
	router.Handle("/unit/update", http.HandlerFunc(unitcontroller.UpdateUnit)).Methods(http.MethodPost)
	router.Handle("/unit/getbyid/{id}", http.HandlerFunc(unitcontroller.GetUnitbyid)).Methods(http.MethodGet)
	router.Handle("/unit/getall", http.HandlerFunc(unitcontroller.GetAllUnits)).Methods(http.MethodGet)

	return router
}
