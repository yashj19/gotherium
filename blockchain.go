package main

import (
	"encoding/json"
	"fmt"
)

var GENESIS_BLOCK_POINTER *Block = nil
var GenesisBlock *Block = NewBlock(GENESIS_BLOCK_POINTER, make(map[string]*Transaction))

type BlockChain struct {
	Blocks []*Block
	// create map to store transactions
	TransactionMap map[string]*Transaction
	CurrentState *GlobalState
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks:         []*Block{},
		TransactionMap: make(map[string]*Transaction),
		CurrentState: NewGlobalState(make(map[string]*Account)),
	}
}

func (bc *BlockChain) EnqueueTransaction(t *Transaction) {
	bc.TransactionMap[Hash(t)] = t
}

func (bc *BlockChain) MineBlock() *Block {
	// create a new block without end state signature
	var block *Block = NewBlock(bc.Blocks[len(bc.Blocks)-1], bc.TransactionMap)

	// get end state signature and update block signature (TODO: probably should make immutable though)
	endStateSignature := bc.GetEndStateSignature(block)
	block.EndStateSignature = endStateSignature
	bc.AddBlock(block)
	bc.TransactionMap = make(map[string]*Transaction) // clear transaction map
	return block
}

func (bc *BlockChain) GetEndState() string {
	return bc.Blocks[len(bc.Blocks)-1].EndStateSignature
}

func (bc *BlockChain) AddBlock(block *Block) {
	if bc.IsBlockValid(block) {
		bc.Blocks = append(bc.Blocks, block)
	} else {
		fmt.Println("Block is not valid")
	}
}

func (bc *BlockChain) IsBlockValid(block *Block) bool {
	// TODO: verify block (transactions, end state signature, previous block) is valid
	return true
}

func (bc *BlockChain) GetEndStateSignature(block *Block) string {
	transactions := block.Transactions
	for _, transaction := range transactions {
		bc.ApplyTransaction(transaction)
	}
	return bc.CurrentState.Signature()
}

func (bc *BlockChain) ApplyTransaction(transaction *Transaction) {
	// TODO: make state immutable later
	senderAccount := bc.CurrentState.Accounts[transaction.SenderAddress]
	if transaction.Nonce != senderAccount.Nonce {
		fmt.Println("Transaction nonce must match sender account nonce")
		return
	}
	senderAccount.Nonce++

	// TODO: verify transaction is well-formed
	if transaction.Amount > senderAccount.Balance {
		fmt.Println("Transaction amount is greater than sender account balance")
		return
	}

	if transaction.RecipientAddress == "" {
		if transaction.Data == "" {
			fmt.Println("Contract creation must give code through the data field")
			return
		} 
		newContractAccount := NewAccount(HashString(transaction.Data), 0, 0, string(transaction.Data), make(map[string]int))
		bc.CurrentState.Accounts[newContractAccount.Address] = newContractAccount
	}

	// if sending to account that doesn't exist yet, create it
	if _, ok := bc.CurrentState.Accounts[transaction.RecipientAddress]; !ok {
		bc.CurrentState.Accounts[transaction.RecipientAddress] = NewAccount(transaction.RecipientAddress, 0, 0, "", make(map[string]int))	
	}

	recipientAccount := bc.CurrentState.Accounts[transaction.RecipientAddress]
	senderAccount.Balance -= transaction.Amount
	recipientAccount.Balance += transaction.Amount

	// execute the contract if account is a contract
	if recipientAccount.Type() == "contract" {
		var args map[string]interface{}
		if err := json.Unmarshal([]byte(transaction.Data), &args); err != nil {
			fmt.Printf("Error parsing contract args: %v\n", err)
			return
		}
		recipientAccount.RunContract(args)
	}

	
}
