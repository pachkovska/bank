package person

import (
	"fmt"
	"github.com/CodingSquire/bank/pkg/money"
)

type client struct {
	ID        int
	FirstName string
	SecName   string
	Cashe     money.Money
}

type bankForClient interface {
	GetBalance(id int) (data money.Money, err error)
	AddToBell(id int, data money.Money) (err error)
	DeductFromBell(id int, data money.Money) (err error)
	CreateAcc(id int) error
}

func (c *client) DeductFromBell(bank bankForClient, change money.Money) (err error) {
	if err = bank.DeductFromBell(c.ID, change); err != nil {
		return fmt.Errorf("error while try to DeductFromBell in bank: %v", err)
	} else {
		c.Cashe = c.Cashe + change
	}
	return
}

func (c client) ShowInfo() {
	fmt.Println("ClientInfo:")
	fmt.Println("ID:", c.ID)
	fmt.Println("FirstName:", c.FirstName)
	fmt.Println("SecName:", c.SecName)
	fmt.Println("Cashe:", c.Cashe)
}

func (c *client) AddToBell(bank bankForClient, balance money.Money) (err error) {
	if c.Cashe-balance < 0 {
		return fmt.Errorf("not enought money")
	}
	c.Cashe = c.Cashe - balance
	err = bank.AddToBell(c.ID, balance)
	if err != nil {
		fmt.Println("error while try to AddToBell in bank")
		return
	}
	fmt.Println("success while try to AddToBell in bank")
	return
}

type Client interface {
	AddToBell(bank bankForClient, balance money.Money) (err error)
	DeductFromBell(bank bankForClient, change money.Money) (err error)
	ShowInfo()
}

func NewClient(id int, firstName, secName string, cashe money.Money) Client {
	return &client{
		ID:        id,
		FirstName: firstName,
		SecName:   secName,
		Cashe:     cashe,
	}
}
