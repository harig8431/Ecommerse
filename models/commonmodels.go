package models


type Normalresponse struct {
	Statuscode  int64  `json:"statuscode"`
	Result      bool   `json:"Result"`
	Description string `json:"desc"`
}
