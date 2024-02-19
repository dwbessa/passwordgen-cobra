package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		symbolsFlag, _ := cmd.Flags().GetBool("symbols")
		numbersFlag, _ := cmd.Flags().GetBool("numbers")
		uppercaseFlag, _ := cmd.Flags().GetBool("uppercase")
		lowercaseFlag, _ := cmd.Flags().GetBool("lowercase")

		var charSet []rune
		if symbolsFlag {
			charSet = append(charSet, symbols...)
		}
		if numbersFlag {
			charSet = append(charSet, numbers...)
		}
		if uppercaseFlag {
			charSet = append(charSet, letters...)
		}
		if lowercaseFlag {
			charSet = append(charSet, lowercaseRunes(letters)...)
		}

		if len(charSet) == 0 {
			fmt.Println("Error: Please choose at least one character type.")
			return
		}

		// Ask for network name and email
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter network name: ")
		networkName, _ := reader.ReadString('\n')
		networkName = strings.TrimSpace(networkName)
		for networkExists(networkName) {
			fmt.Println("Name already in use, use other.")
			networkName, _ = reader.ReadString('\n')
			networkName = strings.TrimSpace(networkName)
		}
		fmt.Print("Enter email: ")
		email, _ := reader.ReadString('\n')
		var password string
		for {
			password := generateRandomString(length, charSet)
			fmt.Println("Generated password:", password)
			fmt.Println("Save this password? (Y/n)")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.ToLower(strings.TrimSpace(response))
			for {
				if response == "y" {
					break
				} else {
					fmt.Println("Generating new password . . .")
					break
				}
			}
			if response == "y" {
				break
			}
		}

		// Create an Account struct
		newAccount := Account{
			Network:  map[string]string{"network": strings.TrimSpace(networkName)},
			Email:    map[string]string{"email": strings.TrimSpace(email)},
			Password: map[string]string{"password": password},
		}

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

		// Append the new account to the slice
		accounts = append(accounts, newAccount)

		// Marshal the updated slice back into JSON
		jsonData, err := json.MarshalIndent(accounts, "", "    ")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}

		// Write the JSON data to the file
		err = os.WriteFile("account.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return
		}
	},
}
