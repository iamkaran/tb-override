/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package theme

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
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

		cssProperty := core.CSSProperty{
			Name:  variableName,
			Value: value,
		}

		err = edit.SetVariable(cfg, themeName, cssProperty)
		if err != nil {
			return err
		}

		return nil
	},
}

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a CSS property",
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

		err = edit.DeleteVariable(cfg, themeName, variableName)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	ThemeCmd.AddCommand(editCmd)
	editCmd.AddCommand(rmCmd)
	editCmd.Flags().StringP("theme", "t", "", "Theme")
	editCmd.Flags().StringP("variable", "", "", "Variable name")
	editCmd.Flags().StringP("value", "", "", "Variable value")
	_ = editCmd.MarkFlagRequired("variable")
	_ = editCmd.MarkFlagRequired("theme")
	_ = editCmd.MarkFlagRequired("value")
	rmCmd.Flags().StringP("theme", "t", "", "Theme")
	rmCmd.Flags().StringP("variable", "", "", "Variable name")
}
