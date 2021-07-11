package main

import (
	"fmt"
	"os"
	//"time"
	"github.com/Iliaromanov/Golang-Blockchain/cmd"
	"github.com/Iliaromanov/Golang-Blockchain/database"
)

func main() {
	cmd.Execute()

	
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Println("Fuck")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()
	/*
	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Transaction{
			database.NewTransaction("ilia", "bob", 5, ""),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()
	fmt.Printf("%s\n", block0hash)
	*/
}