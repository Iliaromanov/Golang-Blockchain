package database

type Account string

func newAccount(value string) Account {
	return Account(value)
}

type Transaction struct {
	From  Account `json:"from"` // JSON encoding struct tag to ensure field name 'From' in Json output
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`	
}

// Creates a new transaction
func newTransaction(from Account, to Account, value uint, data string) Transaction {
	return Transaction{from, to, value, data}
}

// Checks if the transaction was a reward
// (creator of blockchain {me} gets a weekly reward amount of tokens for my brilliant idea)
func isReward(t Transaction) bool {
	return t.Data == "reward"
}
