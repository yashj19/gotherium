package main

import "fmt"

type Transaction struct {
	SenderAddress string
	RecipientAddress string
	Amount int
	Data any
	Nonce int
}

func NewTransaction(senderAddress string, recipientAddress string, amount int, data any, nonce int) *Transaction {
	return &Transaction{
		SenderAddress: senderAddress,
		RecipientAddress: recipientAddress,
		Amount: amount,
		Data: data,
		Nonce: nonce,
	}
}

func (t *Transaction) String() string {
	return fmt.Sprintf("SenderAddress: %s, RecipientAddress: %s, Amount: %d, Data: %v, Nonce: %d", t.SenderAddress, t.RecipientAddress, t.Amount, t.Data, t.Nonce)
}

