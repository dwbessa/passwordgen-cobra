package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")
var symbols = []rune("!@#$%^&*()_+=-[]\\{}|;':\",./<>?")

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

var rootCmd = &cobra.Command{
	Use:   "passwordgen",
	Short: "Generate secure random passwords",
	Long: `Generate secure random passwords according to your desired length 
	and character types.`,
}

type Account struct {
	Network  map[string]string `json:"network"`
	Email    map[string]string `json:"email"`
	Password map[string]string `json:"password"`
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(searchCmd)
	generateCmd.Flags().IntP("length", "l", 12, "Password length")
	generateCmd.Flags().BoolP("symbols", "s", true, "Include symbols")
	generateCmd.Flags().BoolP("numbers", "n", true, "Include numbers")
	generateCmd.Flags().BoolP("uppercase", "u", true, "Include uppercase letters")
	generateCmd.Flags().BoolP("lowercase", "d", true, "Include lowercase letters")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
