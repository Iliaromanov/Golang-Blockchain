package main

import (
	//"log"

	"github.com/Iliaromanov/Golang-Blockchain/cmd"
	//"github.com/Iliaromanov/Golang-Blockchain/database"
)

func main() {
	/*
	// For testing
	s, err := database.NewStateFromDisk()
	if err != nil {
		log.Fatal(err)
	}
	tx := database.Transaction{"ilia", "bob", 5, ""}
	*/

	cmd.Execute()
}