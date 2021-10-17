package bank

import (
	"fmt"
	"github.com/CodingSquire/bank/pkg/api"

	"github.com/CodingSquire/bank/pkg/money"
)

type bank struct {
	Data map[int]money.Money
}

type Bank interface {
	GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error)
	AddToBell(request api.AddToBellRequest) (response api.AddToBellResponse, err error)
	DeductFromBell(request api.DeductFromBellRequest) (response api.DeductFromBellResponse, err error)
	CreateAcc(request api.CreateAccRequest) (response api.CreateAccResponse, err error)
}

func (b *bank) setBalance(id int, data money.Money) {
	b.Data[id] = data
	return
}

func (b *bank) getBalance(id int) (data money.Money, err error) {
	val, ok := b.Data[id]
	if ok {
		data = val

	} else {
		err = fmt.Errorf("no client")
	}
	return
}

func (b *bank) GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error) {
	val, err := b.getBalance(request.ID)
	if err != nil {
		return
	} else {
		response.Data = val
	}
	return
}

func (b *bank) AddToBell(request api.AddToBellRequest) (response api.AddToBellResponse, err error) {
	balance, err := b.getBalance(request.ID)
	if err != nil {
		return response, fmt.Errorf("error while AddToBell:%v", err)
	}
	newBalance := balance + request.Data
	b.setBalance(request.ID, newBalance)
	response.Status = true
	return
}

func (b *bank) DeductFromBell(request api.DeductFromBellRequest) (response api.DeductFromBellResponse, err error) {
	balance, err := b.getBalance(request.ID)
	if err != nil {
		return response, fmt.Errorf("error while DeductFromBell:%v", err)
	}
	newBalance := balance - request.Data
	err = money.ValidateBalance(newBalance)
	if err != nil {
		return response, fmt.Errorf("error while DeductFromBell:%v", err)
	}
	b.setBalance(request.ID, newBalance)
	response.Status = true
	return
}

func (b *bank) CreateAcc(request api.CreateAccRequest) (response api.CreateAccResponse, err error) {
	if _, err := b.getBalance(request.ID); err == nil {
		return response, fmt.Errorf("account reserved")
	}

	b.Data[request.ID] = 0
	response.Status = true
	return
}

func NewBank() Bank {
	return &bank{
		Data: make(map[int]money.Money),
	}
}
