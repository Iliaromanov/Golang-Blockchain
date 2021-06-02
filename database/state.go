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

	// Updating genesis State balances by sequentially 
	//  replaying all database events from transactions.db
	transactionDbFilePath := filepath.Join(cwd, "database", "transaction.db")
	file, err := os.OpenFile(transactionDbFilePath, os.O_APPEND|os.O_RDWR, 0600) // 0600 is readable+writable permission
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	state := &State{balances, make([]Transaction, 0), file}

	// Iterate of transaction.db line by line
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var transaction Transaction
		json.Unmarshal(scanner.Bytes(), &transaction) // parses state.json line into transaction

		// Add transaction to state
		if err := state.apply(transaction); err != nil { // should update balances map and append to transactions slice
			return nil, err
		}
	}

	return state, nil
}

// Add method for State
// adds new transaction to transactionMempool
func (s *State) Add(transaction Transaction) error {
	if err := s.apply(transaction); err != nil {
		return err
	}

	s.transactionMempool = append(s.transactionMempool, transaction)

	return nil
}

// Persist to disk method for State
// writes transactions in transactionMempool to the transaction.db file
func (s *State) Persist() error {
	// make copy of mempool because s.transactinMempool will be modified in loop
	mempool := make([]Transaction, len(s.transactionMempool))
	copy(mempool, s.transactionMempool)

	for _, tx := range mempool {
		txJson, err := json.Marshal(tx)
		if err != nil {
			return err
		}

		if _, err := s.dbFile.Write(append(txJson, '\n')); err != nil {
			return err
		}

		// Remove transaction written to file from mempool
		s.transactionMempool = s.transactionMempool[1:]
	}
	
	return nil
}