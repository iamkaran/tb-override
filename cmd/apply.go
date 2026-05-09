/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"github.com/iamkaran/tb-override/internal/apply"
	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/logger"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a certain theme to make it active",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		log := logger.FromContext(cmd.Context())
		cfg := config.FromContext(cmd.Context())

		err = apply.ApplyTheme(log, cfg, name)
		if err != nil {
			log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringP("name", "n", "", "Name of the theme")
	if err := applyCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
}
