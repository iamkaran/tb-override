/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/logger"
	"github.com/iamkaran/tb-override/internal/setup"
	"github.com/spf13/cobra"
)

// setupCmd represents the install command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the required directories and files required for tb-override to work",
	RunE: func(cmd *cobra.Command, args []string) error {
		log := logger.FromContext(cmd.Context())
		cfg := config.FromContext(cmd.Context())

		err := setup.Setup(cmd.Context(), log, cfg)
		if err != nil {
			if errors.Is(err, core.ErrNoRootPrivilages) {
				return err
			} else {
				return err
			}
		} else {
			fmt.Println("Required directories and files have been created")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
