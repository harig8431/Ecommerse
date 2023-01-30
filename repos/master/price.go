package master

import (
	"fmt"
	"log"
	"project/models"
	"project/utls"
)

type PriceInterface interface {
	CreatePrice(pdtobj *models.Prices) bool
	UpdatePrice(updobj *models.Prices) (models.Prices, bool)
	GetPricebyid(getobj *models.Prices) (models.Prices, bool)
	GetAllPrices() ([]models.Prices, bool)
}
type PriceStruct struct {
}

func (u *PriceStruct) CreatePrice(pricobj *models.Prices) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "price"(
	Price,       
	Date,     
	ProductId,      
	)values($1,$2,$3,$4)RETURNING id`,

		pricobj.Price,
		utls.GetCurrentDateTime(),
		pricobj.ProductId).Scan(&pricobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}
func (u *PriceStruct) UpdatePrice(updobj *models.Prices) (models.Prices, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "price" SET Price=$2,       
	Date=$3,     
	ProductId=$4,     
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Price,
		&updobj.Date,
		&updobj.ProductId,
	)
	//x:=*updobj
	if err != nil {
		fmt.Println("Error in update row in db", err)
		return *updobj, false
	}
	return *updobj, true

}

func (u *PriceStruct) GetPricebyid(getObj *models.Prices) (models.Prices, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Prices{}
	query, _ := MyDb.Prepare(`SELECT Id,Price,Date,ProductId FROM "price" WHERE id=$1`)
	er := query.QueryRow(getObj.Id).Scan(&result.Id,
		&result.Price,
		&result.Date,
		&result.ProductId,
		// &result.Token,
		// hdhgfh
	)

	//x:=*updobj
	if er != nil {
		fmt.Println("Error in finding row in db", er)
		return result, false
	}
	return result, true

}
func (u *PriceStruct) GetAllPrices() ([]models.Prices, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Prices{}
	finalstruct := []models.Prices{}

	query, err := MyDb.Query(`SELECT Id,Price,Date,ProductId FROM "price" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,
			&result.Price,
			&result.Date,
			&result.ProductId,
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
