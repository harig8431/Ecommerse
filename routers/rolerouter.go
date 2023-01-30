package routers

import (
	"net/http"
	"project/Controllers/masters"

	"github.com/gorilla/mux"
)

func RoleRoutes(router *mux.Router) *mux.Router {
	Rolecontroller := masters.Rolecontroller{}
	router.Handle("/Role/create", http.HandlerFunc(Rolecontroller.CreateRole)).Methods(http.MethodPost)
	router.Handle("/Role/update", http.HandlerFunc(Rolecontroller.UpdateRole)).Methods(http.MethodPost)
	router.Handle("/Role/getbyid/{id}", http.HandlerFunc(Rolecontroller.GetRolebyid)).Methods(http.MethodGet)
	router.Handle("/Role/getall", http.HandlerFunc(Rolecontroller.GetAllRoles)).Methods(http.MethodGet)

	return router
}
