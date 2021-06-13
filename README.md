# Blockchain
Simple Blockchain ledger implementation in Go. 

This is a token based ledger and I have decided to name the tokens BLT (Blockchain Ledger Tokens) because it sounds cool.


## CLI Usage

To list balances in current database state execute:
   
> `go run .\main.go balances list`


<!--Hidden Notes:
    * Event-based architecture: production, consumtion, reaction to events (eg. transaction is event, update state is reaction)
    * Reward: for every specific time interval like every week, creator of blockchain gets rewarded a specific amount of tokens like 100
-->
