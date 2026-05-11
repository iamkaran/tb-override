// Package theme contains the declaration of all the sub-commands related to actions performed on themes
/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package theme

import (
	"github.com/spf13/cobra"
)

// themeCmd represents the theme command
var (
	ThemeCmd = &cobra.Command{
		Use:   "theme",
		Short: "Perform actions on themes",
	}
)
