package main

import "fmt"

type Block struct {
	PreviousBlock *Block
	Transactions map[string]*Transaction
	EndStateSignature string
}

func NewBlockWithEndState(previousBlock *Block, transactions map[string]*Transaction, endStateSignature string) *Block {
	return &Block{
		PreviousBlock: previousBlock,
		Transactions: transactions,
		EndStateSignature: endStateSignature,
	}
}

func NewBlock(previousBlock *Block, transactions map[string]*Transaction) *Block {
	return &Block{
		PreviousBlock: previousBlock,
		Transactions: transactions,
		EndStateSignature: "",
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("PreviousBlock: %s, Transactions: %v, EndStateSignature: %s", b.PreviousBlock, b.Transactions, b.EndStateSignature)
}
