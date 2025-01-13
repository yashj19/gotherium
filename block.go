package main

import "fmt"

type Block struct {
	PreviousBlock *Block
	Transactions []*Transaction
	EndStateSignature string
}

func NewBlock(previousBlock *Block, transactions []*Transaction, endStateSignature string) *Block {
	return &Block{
		PreviousBlock: previousBlock,
		Transactions: transactions,
		EndStateSignature: endStateSignature,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("PreviousBlock: %s, Transactions: %v, EndStateSignature: %s", b.PreviousBlock, b.Transactions, b.EndStateSignature)
}
