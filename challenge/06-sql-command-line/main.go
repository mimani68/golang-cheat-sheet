package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yourcli",
	Short: "A brief description of your CLI",
	Long:  `A longer description that spans multiple lines and provides more context.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from root command!")
	},
}

func main() {
	fmt.Println("Hi")
	rootCmd.Execute()
}
