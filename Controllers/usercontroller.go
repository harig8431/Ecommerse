package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/models"
	"project/repos"
	"strconv"

	"github.com/gorilla/mux"
)

type Usercontroller struct {
}

func (u *Usercontroller) Createuser(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := repos.UserInterface(&repos.UserStruct{})
	status := myRepo.CreateUser(&req)
	response := models.Normalresponse{
		Statuscode:  200,
		Result:      status,
		Description: "User Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Usercontroller) Updateuser(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := repos.UserInterface(&repos.UserStruct{})
	value, status := myRepo.UpdateUser(&req)
	response := models.ResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "User Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Usercontroller) Getbyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	userid := models.User{
		Id: id,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := repos.UserInterface(&repos.UserStruct{})
	value, status := myRepo.Getbyid(&userid)
	response := models.ResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "User Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Usercontroller) GetAll(w http.ResponseWriter, r *http.Request) {
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
	myRepo := repos.UserInterface(&repos.UserStruct{})
	value, status := myRepo.GetAll()
	response := models.GetAllUserResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Description: "Users Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
