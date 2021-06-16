package database

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"os"
	"path/filepath"
	"crypto/sha256"
)


type Snapshot [32]byte // Struct to store 32byte hash

type State struct {
	Balances map[Account]uint
	transactionMempool []Transaction
	dbFile *os.File
	snapshot Snapshot // unique hash for latest state update
}

func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd() // gets path to current directory
	if err != nil {
		return nil, err
	}

	genFilePath := filepath.Join(cwd, "database", "genesis.json") // concatenates filepath, adding '/' where necessary
	gen, err := loadGenesis(genFilePath) // load genesis.json into genesis struct
	if err != nil {
		return nil, err
	}

	// Store balances from genesis struct in map
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
	state := &State{balances, make([]Transaction, 0), file, Snapshot{}}

	// Iterate of transaction.db line by line
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var transaction Transaction
		json.Unmarshal(scanner.Bytes(), &transaction) // parse state.json line into transaction struct

		// Add transaction to state
		if err := state.apply(transaction); err != nil { // update balances map
			return nil, err
		}
	}

	err = state.doSnapshot()
	if err != nil {
		return nil, err
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
func (s *State) Persist() (Snapshot, error) {
	// make copy of mempool because s.transactinMempool will be modified in loop
	mempool := make([]Transaction, len(s.transactionMempool))
	copy(mempool, s.transactionMempool)

	for _, tx := range mempool {
		txJson, err := json.Marshal(tx)
		if err != nil {
			return Snapshot{}, err
		}
		
		fmt.Printf("Persisting new Transaction to disk:\n")
		fmt.Printf("\t%s\n", txJson)
		if _, err := s.dbFile.Write(append(txJson, '\n')); err != nil {
			return Snapshot{}, err
		}

		// Create snapshot (hash state after txJson is appended)
		err = s.doSnapshot()
		if err != nil {
			return Snapshot{}, err
		}
		fmt.Printf("New DB Snapshot: %x\n", s.snapshot)

		// Remove transaction written to transaction.db from mempool
		s.transactionMempool = s.transactionMempool[1:]
	}
	
	return s.snapshot, nil
}

func (s *State) Close() {
	s.dbFile.Close()
}

// Apply method for state
// Applies a transaction to state accordingly
func (s *State) apply(tx Transaction) error {
	// If its a reward to creator just add the amount
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	// Check that sender has enough tokens to make transaction
	if tx.Value > s.Balances[tx.From] {
		return fmt.Errorf("insufficient balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value
	return nil
}

func (s *State) doSnapshot() error {
	_, err := s.dbFile.Seek(0, 0) // offsets the reader to start of file
	if err != nil {
		return err
	}

	txsData, err := ioutil.ReadAll(s.dbFile)
	if err != nil {
		return err
	}
	s.snapshot =sha256.Sum256(txsData) // hash the transaction data retrieved from transaction.db

	return nil
}