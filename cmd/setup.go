/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/logger"
	"github.com/iamkaran/tb-override/internal/setup"

	"github.com/spf13/cobra"
)

// setupCmd represents the install command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the required directories and files required for tb-override to work",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.FromContext(cmd.Context())
		cfg := config.FromContext(cmd.Context())
		err := setup.Setup(cmd.Context(), log, cfg)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
