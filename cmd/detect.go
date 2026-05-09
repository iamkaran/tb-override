/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/iamkaran/tb-override/internal/detect"

	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detects the tools necessary for tb-override to work",
	RunE: func(cmd *cobra.Command, args []string) error {
		platform, err := detect.PlatformInfo()
		if err != nil {
			return err
		}

		fmt.Printf("Proxy: %s\n", platform.Proxy.Type)
		fmt.Printf("Supported: %v\n", platform.Proxy.Supported)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}
