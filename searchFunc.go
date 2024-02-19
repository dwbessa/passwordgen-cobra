package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func networkExists(networkName string) bool {
	// Read the existing data from the file
	fileData, err := os.ReadFile("account.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false
	}

	// Unmarshal the JSON data into a slice of Accounts
	var accounts []Account
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &accounts)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return false
		}
	}
	for _, account := range accounts {
		if account.Network["network"] == networkName {
			return true
		}
	}
	return false
}
