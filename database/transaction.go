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

func newTransaction(from Account, to Account, value uint, data string) Transaction {
	return Transaction{from, to, value, data}
}

func isReward(t Transaction) bool {
	return t.Data == "reward"
}
