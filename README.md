# Golang-Blockchain
Blockchain ledger implementation in Go.

*This project is a work in progress*

![Go](https://img.shields.io/badge/-Go-05122A?style=flat&logo=go)&nbsp;

<!-- ![Typescript](https://img.shields.io/badge/-Typescript-05122A?style=flat&logo=typescript)&nbsp;
![HTML](https://img.shields.io/badge/-HTML-05122A?style=flat&logo=HTML5)&nbsp;
![CSS](https://img.shields.io/badge/-CSS-05122A?style=flat&logo=CSS3&logoColor=1572B6)&nbsp;

![Node.js](https://img.shields.io/badge/-Node.js-05122A?style=flat&logo=node.js&logoColor=339933)&nbsp;

![Git](https://img.shields.io/badge/-Git-05122A?style=flat&logo=git)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;

![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
&nbsp; -->

## CLI Usage

For info on available commands and flags execute:

> `go run .\main.go --help`

To list balances in current database state execute:
   
> `go run .\main.go balances list`

To add transaction to ledger execute:
> `go run .\main.go tx add --from={Account sending tokens} --to={Recipients account} --value={Token amount}`

<!--Hidden Notes:
    * Event-based architecture: production, consumtion, reaction to events (eg. transaction is event, update state is reaction)
    * Reward: for every specific time interval like every week, creator of blockchain gets rewarded a specific amount of tokens like 100.
              Balance verification is skipped for reward transactions.
      |-> The balance of the Account who mined a block increases out of the blue as a subject of total tokens supply inflation affecting the whole chain.
    * Blockchain is a database. The token supply, initial user balances, and global blockchain settings are defined in a Genesis file. The Genesis balances indicate what the 
       original blockchain state was and are never updated afterwards. The database state changes are called Transactions (TX).
    * Decentralized: every user has a copy of the blockchain on their disk, so one user can't change the entire blockchain.
    * Batch processing: is the running of "jobs that can run without end user interaction, or can be scheduled to run as resources permit."
      |-> Batching is a common strategy when working with SQL/NoSQL/Other database systems. The batch strategy consist of “handling multiple items at once”. 
          The solution is to   encapsulate transactions to linked “chunks”, “blocks”.
    * Hashing:
         The ParentHash is being used as a reliable “checkpoint,” representing and referencing the previously hashed database content.
         ParentHash improves performance; Only new data + reference to previous state needs to be hashed to achieve immutability.
         E.g., If you attempt to modify a TX value in Block 0, it will result in a new unique Block 0 hash. Hash of Block 1, based on the parent
         Block 0 reference, would therefore immediately change as well. The cascade effect would affect all the blocks, making the malicious
         attacker database invalid - different from the rest of the honest database stakeholders.
         The attacker database would be, therefore, excluded from participating in the network.
    * The time being used for block timestamps is Unix time:
         Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC. The result doesn't depend on the location associated with t. Unix-like
         operating systems often record time as a 32bit count of seconds, but since this method returns a 64bit value it is valid for billions of years into the past or future.

-->

<!--
Blockchain ledger implementation in Go.

This reporsitory contains two separate implementations of a Blockchain in Golang:

The first stores the block information in a .db file database on disk and is not synced for multiple users; contained in DiskBasedBlockchain dir; Visit this dir for usage README.

The second implementation is peer distributed and currently under development.
-->
