package models

type Invoice struct {
	Id         int64      `json:"id"`
	Items      []Products `json:"items"`
	BillAmount int64      `json:"billamount"`
	BillId     int64      `json:"billid"`
	CustomerId int64      `json:"customerid"`
	CreatedOn  string     `json:"createdon"`
	CreatedBy  int64      `json:"createdby"`
}

type InvoicebyId struct {
	Id         int64         `json:"id"`
	Items      []ProductsAll `json:"items"`
	BillAmount int64         `json:"billamount"`
	BillId     int64         `json:"billid"`
	CustomerId int64         `json:"customerid"`
	CreatedOn  string        `json:"createdon"`
	CreatedBy  int64         `json:"createdby"`
}

type SaleCommanRespones struct {
	Statuscode  int64   `json:"statuscode"`
	Status      bool    `json:"status"`
	Value       Invoice `json:"invoicedata"`
	Descreption string  `json:"desc"`
}

type GetAllSaleInvoiceResponse struct {
	Statuscode  int64     `json:"statuscode"`
	Status      bool      `json:"result"`
	Value       []Invoice `json:"value"`
	Descreption string    `json:"desc"`
}

type GetAllSaleInvoiceGetByBillIdResponse struct {
	Statuscode  int64       `json:"statuscode"`
	Status      bool        `json:"result"`
	Value       InvoicebyId `json:"value"`
	Descreption string      `json:"desc"`
}

type GetUserReportByDateRange struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type GetAllSaleInvoiceByDateRangeResponse struct {
	Statuscode  int64         `json:"statuscode"`
	Status      bool          `json:"result"`
	Value       []InvoicebyId `json:"data"`
	Descreption string        `json:"desc"`
}
