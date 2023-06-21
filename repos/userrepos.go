package repos

import (
	"fmt"
	"log"
	"project1/models"
	"project1/utls"
)

type UserInterface interface {
	CreateUser(userobj *models.User) bool
	LoginUser(logobj *models.User) (models.User, bool)
	UpdateUser(updobj *models.User) (models.User, bool)
	Getbyid(getobj *models.User) (models.User, bool)
	GetAll() ([]models.User, bool)
}
type UserStruct struct {
}

func (u *UserStruct) CreateUser(userobj *models.User) bool {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	err := MyDb.QueryRow(`INSERT into "ecomuser"(
	Name,       
	Email,     
	Password,   
	Phone,     
	Address,    
	CreatedOn)values($1,$2,$3,$4,$5,$6)RETURNING id`,
		userobj.Name,
		userobj.Email,
		userobj.Password,
		userobj.Phone,
		userobj.Address,
		utls.GetCurrentDateTime()).Scan(&userobj.Id)
	if err != nil {
		fmt.Println("Error Creating row in db", err)
		return false
	}
	return true

}

func (u *UserStruct) LoginUser(logobj *models.User) (models.User, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	userStruct := models.User{}
	query, err := MyDb.Prepare(`SELECT
	Name,       
	Email,       
	Phone,     
	Address,    
	CreatedOn
	from "ecomuser" where Email=$1 and Password=$2`)
	if err != nil {
		log.Println("Error in User Login QueryRow :", err)
		return userStruct, false
	}
	err = query.QueryRow(logobj.Email, logobj.Password).Scan(
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Phone,
		&userStruct.Address,
		&userStruct.CreatedOn,
	)
	if err != nil {
		fmt.Println("Error in User Login row in db", err)
		return userStruct, false
	}
	return userStruct, true

}
func (u *UserStruct) UpdateUser(updobj *models.User) (models.User, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	query := `UPDATE "ecomuser" SET Name=$2,       
	Email=$3,     
	Phone=$4,     
	Address=$5    
	WHERE id=$1`

	_, err := MyDb.Exec(query,
		&updobj.Id,
		&updobj.Name,
		&updobj.Email,
		&updobj.Phone,
		&updobj.Address,
	)
	//x:=*updobj
	if err != nil {
		fmt.Println("Error in update row in db", err)
		return *updobj, false
	}
	return *updobj, true

}

func (u *UserStruct) Getbyid(getObj *models.User) (models.User, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.User{}
	query, _ := MyDb.Prepare(`SELECT Id, Name,Email,Password,Phone,Address,CreatedOn FROM "ecomuser" WHERE id=$1`)
	er := query.QueryRow(getObj.Id).Scan(&result.Id,
		&result.Name,
		&result.Email,
		&result.Password,
		&result.Phone,
		&result.Address,
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
func (u *UserStruct) GetAll() ([]models.User, bool) {
	MyDb, isconnected := utls.OpenDbConnection()
	if !isconnected {
		fmt.Println("Db connection Failed")
	}
	result := models.User{}
	finalstruct := []models.User{}

	query, err := MyDb.Query(`SELECT Id, Name,Email,Password,Phone,Address,CreatedOn FROM "ecomuser" `)

	if err != nil {
		log.Println("Error in query")
	}
	for query.Next() {
		er := query.Scan(&result.Id,
			&result.Name,
			&result.Email,
			&result.Password,
			&result.Phone,
			&result.Address,
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
