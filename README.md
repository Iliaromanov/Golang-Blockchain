# Golang-Blockchain
Simple Blockchain ledger implementation in Go. 

This is a token based ledger and I have decided to name the tokens BLT (Blockchain Ledger Tokens) because it sounds cool.


## CLI Usage

For info on available commands and flags execute:

> `go run .\main.go --help`

To list balances in current database state execute:
   
> `go run .\main.go balances list`

To add transaction to ledger execute:
> `go run .\main.go tx add --from={Account sending tokens} --to={Account recieving tokens} --value={Amount of tokens sent}`

<!--Hidden Notes:
    * Event-based architecture: production, consumtion, reaction to events (eg. transaction is event, update state is reaction)
    * Reward: for every specific time interval like every week, creator of blockchain gets rewarded a specific amount of tokens like 100
              balance verification is skipped for reward transactions
    * Blockchain is a database. The token supply, initial user balances, and global blockchain settings are defined in a Genesis file. The Genesis balances indicate what the 
       original blockchain state was and are never updated afterwards. The database state changes are called Transactions (TX).
    * Decentralized: every user has a copy of the blockchain on their disk, so one user can't change the entire blockchain.
-->
