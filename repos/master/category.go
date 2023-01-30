package master

import (
	"fmt"
	"log"
	"project/models"
	"project/utls"
)

type CategoryInterface interface {
	CreateCategory(pdtobj *models.Categories) bool
	UpdateCategory(updobj *models.Categories) (models.Categories, bool)
	GetCategorybyid(getobj *models.Categories) (models.Categories, bool)
	GetAllCategories() ([]models.Categories, bool)
}
type Categorystruct struct {
}

func (u *Categorystruct) CreateCategory(prodobj *models.Categories) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "category"(
	Name)values($1)RETURNING id`,
		prodobj.Name,
	).Scan(&prodobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}
func (u *Categorystruct) UpdateCategory(updobj *models.Categories) (models.Categories, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "category" SET Name=$2       
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Name,
	)
	//x:=*updobj
	if err != nil {
		fmt.Println("Error in update row in db", err)
		return *updobj, false
	}
	return *updobj, true

}

func (u *Categorystruct) GetCategorybyid(getObj *models.Categories) (models.Categories, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Categories{}
	query, _ := MyDb.Prepare(`SELECT Id,Name FROM "category" WHERE id=$1`)
	er := query.QueryRow(getObj.Id).Scan(&result.Id,
		&result.Name,

		//&result.Token,
	)

	//x:=*updobj
	if er != nil {
		fmt.Println("Error in finding row in db", er)
		return result, false
	}
	return result, true

}
func (u *Categorystruct) GetAllCategories() ([]models.Categories, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Categories{}
	finalstruct := []models.Categories{}

	query, err := MyDb.Query(`SELECT Id,Name FROM "category" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,
			&result.Name,
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
