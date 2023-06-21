package models

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"Phone"`
	Address   string `json:"address"`
	CreatedOn string `json:"createdon"`
	Token     string `json:"token"`
}

type ResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Result      bool   `json:"Result"`
	Values      User   `json:"Values"`
	Description string `json:"desc"`
}
type GetAllUserResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       []User `json:"value"`
	Description string `json:"desc"`
}
