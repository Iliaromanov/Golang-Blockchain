package database

import (
	"fmt"
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type State struct {
	Balances map[Account]uint
	transactionMempool []Transaction
	dbFile *os.File
	latestBlockHash Hash // unique hash for latest state update
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
	transactionDbFilePath := filepath.Join(cwd, "database", "block.db")
	file, err := os.OpenFile(transactionDbFilePath, os.O_APPEND|os.O_RDWR, 0600) // 0600 is readable+writable permission
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	state := &State{balances, make([]Transaction, 0), file, Hash{}}

	// Iterate of transaction.db line by line
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		blockWrapJson := scanner.Bytes() // read json string
		var blockWrp BlockWrap 
		err = json.Unmarshal(blockWrapJson, &blockWrp) // unmarshal json string into BlockWrap struct
		if err != nil {
			return nil, err
		}

		// Add transaction to state
		err = state.applyBlock(blockWrp.Value)
		if err != nil { // update balances map
			return nil, err
		}
		state.latestBlockHash = blockWrp.Key // update latest block hash value
	}

	return state, nil
}

func (s *State) AddBlock(b Block) error {
	for _, tx := range b.Transactions {
		if err := s.AddTx(tx); err != nil {
			return err
		}
	}

	return nil
}

// Add method for State
// adds new transaction to transactionMempool
func (s *State) AddTx(transaction Transaction) error {
	if err := s.apply(transaction); err != nil {
		return err
	}
	s.transactionMempool = append(s.transactionMempool, transaction)

	return nil
}

// Persist to disk method for State
// writes transactions in transactionMempool to the transaction.db file
func (s *State) Persist() (Hash, error) {
	// Create block
	block := NewBlock(
		s.latestBlockHash,
		uint64(time.Now().Unix()),
		s.transactionMempool,
	)

	// Generate hash for the block
	blockHash, err := block.Hash()
	if err != nil {
		return Hash{}, err
	}

	// Encode block into JSON formatted string
	blockWrp := BlockWrap{blockHash, block}
	blockWrpJson, err := json.Marshal(blockWrp)
	if err != nil {
		return Hash{}, err 
	}

	fmt.Println("Persisting new block to disk:")
	fmt.Printf("\t%s\n", blockWrpJson)

	// Write the block to DB file on disk
	_, err = s.dbFile.Write(append(blockWrpJson, '\n'))
	if err != nil {
		return Hash{}, err
	}

	s.latestBlockHash = blockHash // Update latest block hash value
	s.transactionMempool = []Transaction{} // reset transaction mempool

	return s.latestBlockHash, nil
}

func (s *State) Close() {
	s.dbFile.Close()
}

func (s *State) applyBlock(b Block) error {
	for _, tx := range b.Transactions {
		if err := s.apply(tx); err != nil {
			return err
		}
	}
	return nil
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
