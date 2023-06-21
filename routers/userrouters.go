package routers

import (
	"net/http"
	"project1/Controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router{
usercontroller:=Controllers.Usercontroller{}
router.Handle("/user/createuser",http.HandlerFunc(usercontroller.Createuser)).Methods(http.MethodPost)
router.Handle("/user/login",http.HandlerFunc(usercontroller.LoginUser)).Methods(http.MethodPost)
router.Handle("/user/updateuser",http.HandlerFunc(usercontroller.Updateuser)).Methods(http.MethodPost)
router.Handle("/user/getbyid/{id}",http.HandlerFunc(usercontroller.Getbyid)).Methods(http.MethodGet)
router.Handle("/user/getall",http.HandlerFunc(usercontroller.GetAll)).Methods(http.MethodGet)

return router
}