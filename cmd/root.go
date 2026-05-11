// Package cmd contains the declaration of all the command line flags related to tb-override
/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"os"

	"github.com/iamkaran/tb-override/cmd/list"
	"github.com/iamkaran/tb-override/cmd/theme"
	"github.com/iamkaran/tb-override/internal/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "tb-override",
		Short: "White-labeling for ThingsBoard Community Edition",

		// Runs before any command logic
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := config.InitializeConfig(cfgFile, cmd)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(theme.ThemeCmd)
	rootCmd.AddCommand(list.ListCmd)
}
