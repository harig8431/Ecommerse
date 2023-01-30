package models

type Categories struct {
	Id   string  `json:"id"`
	Name string `json:"name"`
}
type CategoryResponseModel struct {
	Statuscode  int64      `json:"statuscode"`
	Result      bool       `json:"Result"`
	Values      Categories `json:"Values"`
	Description string     `json:"desc"`
}

type GetAllCategoriesResponseModel struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"status"`
	Value       []Categories `json:"value"`
	Description string       `json:"desc"`
}
