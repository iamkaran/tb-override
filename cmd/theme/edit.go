/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package theme

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/edit"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a theme by adding CSS overrides",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.FromContext(cmd.Context())

		themeName, err := cmd.Flags().GetString("theme")
		if err != nil {
			return err
		}

		variableName, err := cmd.Flags().GetString("variable")
		if err != nil {
			return err
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}

		err = edit.AppendVariable(cfg, themeName, variableName, value)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	ThemeCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("theme", "", "", "Theme")
	editCmd.Flags().StringP("variable", "", "", "Variable name")
	editCmd.Flags().StringP("value", "", "", "Variable value")
	_ = editCmd.MarkFlagRequired("variable")
	_ = editCmd.MarkFlagRequired("value")
	_ = editCmd.MarkFlagRequired("theme")
}
