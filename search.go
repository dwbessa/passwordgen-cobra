package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an account by network name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a network name.")
			return
		}
		networkName := args[0]

		// Read the existing data from the file
		fileData, err := os.ReadFile("account.json")
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		// Unmarshal the JSON data into a slice of Accounts
		var accounts []Account
		if len(fileData) > 0 {
			err = json.Unmarshal(fileData, &accounts)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
		}

		// Iterate over the accounts and print the one with the specified network name
		for _, account := range accounts {
			if account.Network["network"] == networkName {
				fmt.Println("Account: ", account.Network["network"])
				fmt.Println("Email: ", account.Email["email"])
				fmt.Println("Password: ", account.Password["password"])
				return
			}
		}

		fmt.Println("No account found with the specified network name.")
	},
}
