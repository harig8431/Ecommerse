package models

type Products struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Quantity  string `json:"quantity"`
	Category  string `json:"category"`
	Unit      string `json:"unit"`
	Price     int64 `json:"price"`
	CreatedOn string `json:"createdon"`
	Token     string `json:"token"`
}
type ProductsAll struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Quantity  string     `json:"quantity"`
	Category  Categories `json:"category"`
	Unit      Unitss     `json:"unit"`
	Price     Prices     `json:"price"`
	CreatedOn string     `json:"createdon"`
	Token     string     `json:"token"`
}
type ProductResponses struct {
	Statuscode  int64       `json:"statuscode"`
	Result      bool        `json:"Result"`
	Values      ProductsAll `json:"Values"`
	Description string      `json:"desc"`
}
type ProdResponseModel struct {
	Statuscode  int64    `json:"statuscode"`
	Result      bool     `json:"Result"`
	Values      Products `json:"Values"`
	Description string   `json:"desc"`
}

type GetAllProductResponseModel struct {
	Statuscode  int64      `json:"statuscode"`
	Status      bool       `json:"status"`
	Value       []Products `json:"value"`
	Description string     `json:"desc"`
}
