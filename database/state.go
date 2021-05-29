package database

import (
	"fmt"
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	Balances map[Account]uint
	transactionMempool []Transaction
	dbFile *os.File
}

func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd() // gets path to current directory
	if err != nil {
		return nil, err
	}

	genFilePath := filepath.Join(cwd, "database", "genesis.json") // concatenates filepath, adding '/' where necessary
	gen, err := loadGenesis(genFilePath) // maybe use os.Open()?
	if err != nil {
		return nil, err
	}

	// Store balances from genesis.json
	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}
	
}