package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(scanRoutesCmd)
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		fmt.Print("Enter new user name: ")
		fmt.Scanln(&name) // Reading user input
		if name != "" {
			fmt.Printf("Hello! %v Nice to meet you.", name)
		}
	},
}
