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

type Pricecontroller struct {
}

func (u *Pricecontroller) CreatePrice(w http.ResponseWriter, r *http.Request) {
	req := models.Prices{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.PriceInterface(&master.PriceStruct{})
	status := myRepo.CreatePrice(&req)
	response := models.PriceResponseModel{
		Statuscode:  200,
		Result:      status,
		Description: "Price Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Pricecontroller) UpdatePrice(w http.ResponseWriter, r *http.Request) {
	req := models.Prices{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.PriceInterface(&master.PriceStruct{})
	value, status := myRepo.UpdatePrice(&req)
	response := models.PriceResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Price Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Pricecontroller) GetPricebyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	Pricestruct := models.Prices{
		Id: id,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := master.PriceInterface(&master.PriceStruct{})
	value, status := myRepo.GetPricebyid(&Pricestruct)
	response := models.PriceResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Price Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Pricecontroller) GetAllPrices(w http.ResponseWriter, r *http.Request) {
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
	myRepo := master.PriceInterface(&master.PriceStruct{})
	value, status := myRepo.GetAllPrices()
	response := models.GetAllPriceResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Description: "Prices Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
