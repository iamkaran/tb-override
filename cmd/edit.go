/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a theme by adding CSS overrides",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("variable", "", "", "Variable name")
	if err := editCmd.MarkFlagRequired("variable"); err != nil {
		panic(err)
	}
	editCmd.Flags().StringP("value", "", "", "Variable value")
	if err := editCmd.MarkFlagRequired("value"); err != nil {
		panic(err)
	}
}
