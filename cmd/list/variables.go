/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package list

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listVariablesCmd = &cobra.Command{
	Use:   "variables",
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
	ListCmd.AddCommand(listVariablesCmd)
	listVariablesCmd.Flags().StringP("by-category", "c", "", "Filter by category")
	listVariablesCmd.Flags().BoolP("list-categories", "l", false, "List all categories available")
	listVariablesCmd.Flags().BoolP("list-all", "a", false, "List all variables")
	listVariablesCmd.MarkFlagsOneRequired("by-category", "list-categories", "list-all")
}
