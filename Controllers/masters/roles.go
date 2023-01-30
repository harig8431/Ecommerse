package masters

import (
	"encoding/json"
	"log"
	"net/http"
	"project/models"
	"project/repos/master"

	"strconv"

	"github.com/gorilla/mux"
)

type Rolecontroller struct {
}

func (u *Rolecontroller) CreateRole(w http.ResponseWriter, r *http.Request) {
	req := models.Roles{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.RoleInterface(&master.RoleStruct{})
	status := myRepo.CreateRole(&req)
	response := models.RoleResponseModel{
		Statuscode:  200,
		Result:      status,
		Description: "Role Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Rolecontroller) UpdateRole(w http.ResponseWriter, r *http.Request) {
	req := models.Roles{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.RoleInterface(&master.RoleStruct{})
	value, status := myRepo.UpdateRole(&req)
	response := models.RoleResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Role Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Rolecontroller) GetRolebyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	Roleid := models.Roles{
		Id: id,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := master.RoleInterface(&master.RoleStruct{})
	value, status := myRepo.GetRolebyid(&Roleid)
	response := models.RoleResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Role Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Rolecontroller) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	//req:=models.User{}
	// request := mux.Vars(r)
	// id, err := strconv.ParseInt(request["id"], 10, 64)
	// userid := models.User{
	// 	Id: id,
	// }
	// if err != nil {
	// 	log.Println("Error in Decoding UserGetById Request :", err)
	// }
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := master.RoleInterface(&master.RoleStruct{})
	value, status := myRepo.GetAllRoles()
	response := models.GetAllRoleResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Description: "Categories Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
