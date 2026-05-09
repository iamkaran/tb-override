/*
Copyright © 2026 Karanveer Singh kforkaranveer@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/iamkaran/tb-override/internal/detect"
	"github.com/iamkaran/tb-override/internal/logger"

	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detects the tools necessary for tb-override to work",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.FromContext(cmd.Context())
		platform, err := detect.PlatformInfo()
		if err != nil {
			log.Error("Couldn't get platform info", "error", err)
		}
		fmt.Printf("Proxy: %s\n", platform.Proxy.Type)
		fmt.Printf("Supported: %v\n", platform.Proxy.Supported)
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}
