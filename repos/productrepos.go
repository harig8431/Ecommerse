package repos

import (
	"fmt"
	"log"
	"project/models"
	"project/utls"
)

type ProductInterface interface {
	CreateProduct(pdtobj *models.Products) bool
	UpdateProduct(updobj *models.Products) (models.Products, bool)
	GetProductbyid(getobj *int64) (models.Products, bool)
	GetAllProducts() ([]models.Products, bool)
}
type ProductStruct struct {
}

func (u *ProductStruct) CreateProduct(prodobj *models.Products) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "products"(
	Name,       
	Quantity,     
	Category,   
	Unit,     
	Price,    
	CreatedOn)values($1,$2,$3,$4,$5,$6)RETURNING id`,
		prodobj.Name,
		prodobj.Quantity,
		prodobj.Category,
		prodobj.Unit,
		prodobj.Price,
		utls.GetCurrentDateTime()).Scan(&prodobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}
func (u *ProductStruct) UpdateProduct(updobj *models.Products) (models.Products, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "products" SET Name=$2,       
	Quantity=$3,     
	Category=$4,     
	Unit=$5,    
	Price=$6
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Name,
		&updobj.Quantity,
		&updobj.Category,
		&updobj.Unit,
		&updobj.Price,
	)
	//x:=*updobj
	if err != nil {
		fmt.Println("Error in update row in db", err)
		return *updobj, false
	}
	return *updobj, true

}

func (u *ProductStruct) GetProductbyid(getObj *int64) (models.Products, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Products{}
	query, _ := MyDb.Prepare(`SELECT Id,Name,Quantity,Category,Unit,Price,CreatedOn FROM "products" WHERE id=$1`)
	er := query.QueryRow(getObj).Scan(&result.Id,
		&result.Name,
		&result.Quantity,
		&result.Category,
		&result.Unit,
		&result.Price,
		&result.CreatedOn,
		//&result.Token,
	)

	//x:=*updobj
	if er != nil {
		fmt.Println("Error in finding row in db", er)
		return result, false
	}
	return result, true

}
func (u *ProductStruct) GetAllProducts() ([]models.Products, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Products{}
	finalstruct := []models.Products{}

	query, err := MyDb.Query(`SELECT Id,Name,Quantity,Category,Unit,Price,CreatedOn FROM "products" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,
			&result.Name,
			&result.Quantity,
			&result.Category,
			&result.Unit,
			&result.Price,
			&result.CreatedOn,
			//&result.Token,
		)
		if er != nil {
			fmt.Println("Error in finding row in db", er)
			return finalstruct, false
		}
		finalstruct = append(finalstruct, result)

	}

	//x:=*updobj

	return finalstruct, true

}
