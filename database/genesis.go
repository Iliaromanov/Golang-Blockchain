package database

import (
	"io/ioutil"
	"encoding/json"
)

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := ioutil.ReadFile(path) // read all content from genesis.json
	if err != nil {
		return genesis{}, err
	}

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis) // parse content into loadedGenesis struct
	if err != nil {
		return genesis{}, err
	}

	return loadedGenesis, nil
}
