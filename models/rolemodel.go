package models

type Roles struct {
	Id        int64  `json:"id"`
	Role     string `json:"Role"`
	Date      string `json:"date"`
	ProductId string `json:"productid"`
}
type RoleResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Result      bool   `json:"Result"`
	Values      Roles `json:"Values"`
	Description string `json:"desc"`
}
type GetAllRoleResponseModel struct {
	Statuscode  int64    `json:"statuscode"`
	Status      bool     `json:"status"`
	Value       []Roles `json:"value"`
	Description string   `json:"desc"`
}
