package models

type Category struct {
	Id   string  `json:"id"`
	Name string `json:"name"`
}
type CategoryResponseModel struct {
	Statuscode  int64      `json:"statuscode"`
	Result      bool       `json:"Result"`
	Values      Category `json:"Values"`
	Description string     `json:"desc"`
}

type GetAllCategoriesResponseModel struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"status"`
	Value       []Category `json:"value"`
	Description string       `json:"desc"`
}
