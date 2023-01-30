package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/models"
	"project/repos"
	"project/repos/master"
	"strconv"

	"github.com/gorilla/mux"
)

type Productcontroller struct {
}

func (u *Productcontroller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	req := models.Products{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := repos.ProductInterface(&repos.ProductStruct{})
	status := myRepo.CreateProduct(&req)
	response := models.ProdResponseModel{
		Statuscode:  200,
		Result:      status,
		Description: "Product Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
func (u *Productcontroller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	req := models.Products{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding ", err)
	}
	myRepo := repos.ProductInterface(&repos.ProductStruct{})
	value, status := myRepo.UpdateProduct(&req)
	response := models.ProdResponseModel{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Product Created Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Productcontroller) GetProductbyid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	// productid := models.Products{
	// 	Id: id,
	// }
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	}
	// err:=json.NewDecoder(r.Body).Decode(&req)
	// if err!=nil {
	// 	log.Println("Error decoding ",err)
	// }
	repo := repos.ProductInterface(&repos.ProductStruct{})

	Value, status := repo.GetProductbyid(&id)
	categoryStruct := models.Categories{
		Id: Value.Category,
	}
	UnitStruct := models.Unitss{
		Id: Value.Unit,
	}
	PriceStruct := models.Prices{
		Id: Value.Price,
	}
	categoryrepo := master.CategoryInterface(&master.Categorystruct{})
	unitrepo := master.UnitInterface(&master.Unitsstruct{})
	pricerepo := master.PriceInterface(&master.PriceStruct{})

	category, _ := categoryrepo.GetCategorybyid(&categoryStruct)
	unit, _ := unitrepo.GetUnitbyid(&UnitStruct)
	price, _ := pricerepo.GetPricebyid(&PriceStruct)

	value := models.ProductsAll{
		Id:        Value.Id,
		Name:      Value.Name,
		Category:  category,
		Quantity:  Value.Quantity,
		Unit:      unit,
		Price:     price,
		CreatedOn: Value.CreatedOn,
	}
	response := models.ProductResponses{
		Statuscode:  200,
		Result:      status,
		Values:      value,
		Description: "Product listedf successful",
	}

	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}

func (u *Productcontroller) GetAllProducts(w http.ResponseWriter, r *http.Request) {
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
	myRepo := repos.ProductInterface(&repos.ProductStruct{})
	value, status := myRepo.GetAllProducts()
	response := models.GetAllProductResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Description: "Products Found Successfully",
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshaling ", err)
	}
	w.Write(resp)
}
