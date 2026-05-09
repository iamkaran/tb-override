/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/create"
	"github.com/iamkaran/tb-override/internal/logger"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new theme and store it in the themes directory",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		log := logger.FromContext(cmd.Context())
		cfg := config.FromContext(cmd.Context())

		err = create.CreateTheme(log, cfg, name)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the theme")
	if err := createCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
}
