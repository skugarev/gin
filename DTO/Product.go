package DTO

type ProductDTO struct {
	Id            uint   `json:"id,string"`
	ValletId      uint   `json:"valletId,string"`
	OperationType string `json:"operationtype,string"`
	Amount        uint   `json:"amount,string"`
}
