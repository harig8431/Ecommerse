package models

type Units struct {
	Id   string `json:"id"`
	Unit string `json:"unit"`
}
type UnitResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Result      bool   `json:"Result"`
	Values      Units  `json:"Values"`
	Description string `json:"desc"`
}
type GetAllUnitsResponseModel struct {
	Statuscode  int64   `json:"statuscode"`
	Status      bool    `json:"status"`
	Value       []Units `json:"value"`
	Description string  `json:"desc"`
}
