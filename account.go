package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Account: struct that represents an account (externally owned or contract) in the blockchain
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

// runs the contract (takes in args parsed from JSON) - TODO: Implement actual code execution engine
// for now, just writes the literal go code to a temp file and runs it (very bad and insecure)
func (a *Account) RunContract(args map[string]interface{}) {
	// create a temporary file to hold the code
	tmpFile, err := os.CreateTemp("", "contract*.go")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	// write the code to the temp file
	_, err = tmpFile.WriteString(fmt.Sprintf(`
package main

import "fmt"

func main() {
	%s
}
`, a.Code))
	if err != nil {
		fmt.Printf("Error writing code: %v\n", err)
		return
	}
	tmpFile.Close()

	// build and run the code
	cmd := exec.Command("go", "run", tmpFile.Name())
	
	// convert args to environment variables
	env := os.Environ()
	for k, v := range args {
		env = append(env, fmt.Sprintf("%s=%v", k, v))
	}
	cmd.Env = env

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing code: %v\n", err)
		fmt.Printf("Output: %s\n", output)
		return
	}

	fmt.Printf("Contract execution output: %s\n", output)


}