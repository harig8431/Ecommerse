package master

import (
	"fmt"
	"log"
	"project1/models"
	"project1/utls"
)

type RoleInterface interface {
	CreateRole(pdtobj *models.Roles) bool
	UpdateRole(updobj *models.Roles) (models.Roles, bool)
	GetRolebyid(getobj *models.Roles) (models.Roles, bool)
	GetAllRoles() ([]models.Roles, bool)
}
type RoleStruct struct {
}

func (u *RoleStruct) CreateRole(pricobj *models.Roles) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "roles"(
	Role,       
	Date,     
	ProductId,      
	)values($1,$2,$3,$4)RETURNING id`,

		pricobj.Role,
		utls.GetCurrentDateTime(),
		pricobj.ProductId).Scan(&pricobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}
func (u *RoleStruct) UpdateRole(updobj *models.Roles) (models.Roles, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "roles" SET Role=$2,       
	Date=$3,     
	ProductId=$4,     
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Role,
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

func (u *RoleStruct) GetRolebyid(getObj *models.Roles) (models.Roles, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Roles{}
	query, _ := MyDb.Prepare(`SELECT Id,Role,Date,ProductIdFROM "roles" WHERE id=$1`)
	er := query.QueryRow(getObj.Id).Scan(&result.Id,
		&result.Role,
		&result.Date,
		&result.ProductId,
		//&result.Token,
	)

	//x:=*updobj
	if er != nil {
		fmt.Println("Error in finding row in db", er)
		return result, false
	}
	return result, true

}
func (u *RoleStruct) GetAllRoles() ([]models.Roles, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.Roles{}
	finalstruct := []models.Roles{}

	query, err := MyDb.Query(`SELECT Id,Role,Date,ProductId FROM "roles" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,
			&result.Role,
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
