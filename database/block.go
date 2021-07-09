package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash [32]byte // Struct to store 32byte hash


// Converts string hash to byte array
func (h Hash) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h[:])), nil
}

// Converts sha256.Sum256 []byte array output to string
func (h *Hash) UnmarshalText(data []byte) error {
	_, err := hex.Decode(h[:], data)
	return err
}


// Main block struct
type Block struct {
	Header       BlockHeader   `json:"header"` // Blocks metadata (parent block hash + time)
	Transactions []Transaction `json:"payload"`// new transactions only (payload)
}

// Struct for storing blocks metadata
type BlockHeader struct {
	Parent Hash   `json:"parent"` // hash of parent block
	Time   uint64 `json:"timestamp"`
}

// Struct for storing a Block and its corresponding hash
type BlockWrap struct {
	Key   Hash  `json:"hash"`
	Value Block `json:"block"`
}

// Creates new block
func NewBlock(parent Hash, time uint64, txs []Transaction) Block {
	return Block{BlockHeader{parent, time}, txs}
}

// Hashes block struct encoded as JSON
func (b Block) Hash() (Hash, error) {
	blockJson, err := json.Marshal(b)

	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(blockJson), nil
}
