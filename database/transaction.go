package database

type Account string

func NewAccount(value string) Account {
	return Account(value)
}

type Transaction struct {
	From  Account `json:"from"` // JSON encoding struct tag to ensure field name 'From' in JSON file
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`	
}

// Creates a new transaction
func NewTransaction(from Account, to Account, value uint, data string) Transaction {
	return Transaction{from, to, value, data}
}

// Checks if the transaction was a reward
// (creator of blockchain {me} gets a weekly reward amount of tokens for my brilliant idea)
func IsReward(t Transaction) bool {
	return t.Data == "reward"
}
