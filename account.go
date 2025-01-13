package main

import (
	"fmt"
)

type Account struct {
	Address string
	Balance int
	Nonce   int
	Code    string
	Storage map[string]int
}

func NewAccount(address string, balance int, nonce int, code string, storage map[string]int) *Account {
	return &Account{
		Address: address,
		Balance: balance,
		Nonce:   nonce,
		Code:    code,
		Storage: storage,
	}
}

func (a *Account) String() string {
	return fmt.Sprintf("Address: %s, Balance: %d, Nonce: %d, Code: %s, Storage: %v", a.Address, a.Balance, a.Nonce, a.Code, a.Storage)
}

func (a *Account) Type() string {
	if a.Code == "" {
		return "contract"
	}
	return "externally owned"
}
