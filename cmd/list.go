/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of common CSS properties of ThingsBoard CE",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.FromContext(cmd.Context())

		err := list.ListVariables(cmd.Context(), cfg, cmd)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("by-category", "c", "", "Filter by category. Run with -l flag to list all categories")
	listCmd.Flags().BoolP("list-categories", "l", false, "List all categories available")
}
