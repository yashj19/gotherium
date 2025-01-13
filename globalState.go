package main

import "fmt"

// GlobalState: struct that represents the global state of the blockchain,
// which is just the current state of all accounts stored on it
type GlobalState struct {
	Accounts map[string]*Account
}

func NewGlobalState(accounts map[string]*Account) *GlobalState {
	return &GlobalState{
		Accounts: accounts,
	}
}

func (gs *GlobalState) String() string {
	return fmt.Sprintf("GlobalState: %v", gs.Accounts)
}

func (gs *GlobalState) Signature() string {
	return gs.String()
}

