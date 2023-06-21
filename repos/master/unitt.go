package master

import (
	"fmt"
	"log"
	"project1/models"
	"project1/utls"
)

type UnitInterface interface {
	CreateUnit(pdtobj *models.Units) bool
	UpdateUnit(updobj *models.Units) (models.Units, bool)
	GetUnitbyid(getobj *models.Units) (models.Units, bool)
	GetAllUnitss() ([]models.Units, bool)
}
type Unitsstruct struct {
}

func (u *Unitsstruct) CreateUnit(prodobj *models.Units) bool {
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
func (u *Unitsstruct) UpdateUnit(updobj *models.Units) (models.Units, bool) {
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

func (u *Unitsstruct) GetUnitbyid(getObj *models.Units) (models.Units, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Units{}
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
func (u *Unitsstruct) GetAllUnitss() ([]models.Units, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Units{}
	finalstruct := []models.Units{}

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
