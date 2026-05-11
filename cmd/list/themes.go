/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package list

import (
	"fmt"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/fs"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "List themes",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.FromContext(cmd.Context())

		listOfThemes, err := fs.ListDirs(filepath.Join(
			cfg.TBOverride.Dirs.RootDirectory,
			cfg.TBOverride.Dirs.ThemesDirectory,
		))

		if err != nil {
			return err
		}

		activeTheme, err := fs.GetActiveTheme(cfg)
		if err != nil {
			return err
		}

		for _, theme := range listOfThemes {
			if theme == activeTheme {
				fmt.Printf("%s *\n", theme)
			} else {
				fmt.Println(theme)
			}
		}

		return nil
	},
}

func init() {
	ListCmd.AddCommand(listThemesCmd)
}
