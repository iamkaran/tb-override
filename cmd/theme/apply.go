/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package theme

import (
	"github.com/iamkaran/tb-override/internal/apply"
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a certain theme to make it active",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		cfg := config.FromContext(cmd.Context())

		err = apply.ApplyTheme(cfg, name)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	ThemeCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringP("name", "n", "", "Name of the theme")
	_ = applyCmd.MarkFlagRequired("name")
}
