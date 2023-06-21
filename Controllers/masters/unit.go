package masters

import (
	"encoding/json"
	"log"
	"net/http"
	"project1/models"
	"project1/repos/master"
	"strconv"

	"github.com/gorilla/mux"
)

type Unitcontroller struct {
}

func (u *Unitcontroller) CreateUnit(w http.ResponseWriter, r *http.Request) {
	req := models.Units{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.UnitInterface(&master.Unitsstruct{})
	status := myRepo.CreateUnit(&req)
	response := models.UnitResponseModel{
		Statuscode:  200,
		Result:      status,
		Description: "Unit Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Unitcontroller) UpdateUnit(w http.ResponseWriter, r *http.Request) {
	req := models.Units{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.UnitInterface(&master.Unitsstruct{})
	value, status := myRepo.UpdateUnit(&req)
	response := models.UnitResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Unit Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Unitcontroller) GetUnitbyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	unitid := strconv.FormatInt(id, 10)
	Unitid := models.Units{
		Id: unitid,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := master.UnitInterface(&master.Unitsstruct{})
	value, status := myRepo.GetUnitbyid(&Unitid)
	response := models.UnitResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Unit Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Unitcontroller) GetAllUnits(w http.ResponseWriter, r *http.Request) {
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
	myRepo := master.UnitInterface(&master.Unitsstruct{})
	value, status := myRepo.GetAllUnitss()
	response := models.GetAllUnitsResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Description: "Units Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
