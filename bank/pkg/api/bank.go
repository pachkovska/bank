package api

import "github.com/CodingSquire/bank/pkg/money"

type GetBalanceRequest struct {
	ID int `json:"id"`
}

type GetBalanceResponse struct {
	Data money.Money `json:"data"`
}

type AddToBellRequest struct {
	ID   int         `json:"id"`
	Data money.Money `json:"data"`
}

type AddToBellResponse struct {
	Status bool `json:"status"`
}

type DeductFromBellRequest struct {
	ID   int         `json:"id"`
	Data money.Money `json:"data"`
}

type DeductFromBellResponse struct {
	Status bool `json:"status"`
}

type CreateAccRequest struct {
	ID int `json:"id"`
}

type CreateAccResponse struct {
	Status bool `json:"status"`
}
