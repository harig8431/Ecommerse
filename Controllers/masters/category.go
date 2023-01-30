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

type Categorycontroller struct {
}

func (u *Categorycontroller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := models.Categories{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.CategoryInterface(&master.Categorystruct{})
	status := myRepo.CreateCategory(&req)
	response := models.CategoryResponseModel{
		Statuscode:  200,
		Result:      status,
		Description: "Category Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Categorycontroller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	req := models.Categories{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := master.CategoryInterface(&master.Categorystruct{})
	value, status := myRepo.UpdateCategory(&req)
	response := models.CategoryResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Category Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Categorycontroller) GetCategorybyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	categoryid := strconv.FormatInt(id, 10)
	Categorystruct := models.Categories{
		Id: categoryid,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	myRepo := master.CategoryInterface(&master.Categorystruct{})
	value, status := myRepo.GetCategorybyid(&Categorystruct)
	response := models.CategoryResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Category Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Categorycontroller) GetAllCategorys(w http.ResponseWriter, r *http.Request) {
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
	myRepo := master.CategoryInterface(&master.Categorystruct{})
	value, status := myRepo.GetAllCategories()
	response := models.GetAllCategoriesResponseModel{
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
