package master

import (
	"fmt"
	"log"
	"project/models"
	"project/utls"
)

type UnitInterface interface {
	CreateUnit(pdtobj *models.Unitss) bool
	UpdateUnit(updobj *models.Unitss) (models.Unitss, bool)
	GetUnitbyid(getobj *models.Unitss) (models.Unitss, bool)
	GetAllUnitss() ([]models.Unitss, bool)
}
type Unitsstruct struct {
}

func (u *Unitsstruct) CreateUnit(prodobj *models.Unitss) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "units"(   
	Unit)values($1)RETURNING id`,
		prodobj.Unit).Scan(&prodobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}
func (u *Unitsstruct) UpdateUnit(updobj *models.Unitss) (models.Unitss, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "units" SET   
	Unit=$2    
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Unit,
	)
	//x:=*updobj
	if err != nil {
		fmt.Println("Error in update row in db", err)
		return *updobj, false
	}
	return *updobj, true

}

func (u *Unitsstruct) GetUnitbyid(getObj *models.Unitss) (models.Unitss, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Unitss{}
	query, _ := MyDb.Prepare(`SELECT Id,Unit FROM "units" WHERE id=$1`)
	er := query.QueryRow(getObj.Id).Scan(&result.Id,
		&result.Unit,
	)

	//x:=*updobj
	if er != nil {
		fmt.Println("Error in finding row in db", er)
		return result, false
	}
	return result, true

}
func (u *Unitsstruct) GetAllUnitss() ([]models.Unitss, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Unitss{}
	finalstruct := []models.Unitss{}

	query, err := MyDb.Query(`SELECT Id,Unit  FROM "units" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,

			&result.Unit,
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
