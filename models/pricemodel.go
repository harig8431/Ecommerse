package models

type Prices struct {
	Id        int64  `json:"id"`
	Price     int64 `json:"price"`
	Date      string `json:"date"`
	ProductId string `json:"productid"`
}
type PriceResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Result      bool   `json:"Result"`
	Values      Prices `json:"Values"`
	Description string `json:"desc"`
}
type GetAllPriceResponseModel struct {
	Statuscode  int64    `json:"statuscode"`
	Status      bool     `json:"status"`
	Value       []Prices `json:"value"`
	Description string   `json:"desc"`
}
