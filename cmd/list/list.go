// Package list contains sub-commands for listing variables and themes
/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package list

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var (
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List of common CSS properties of ThingsBoard CE",
	}
)
