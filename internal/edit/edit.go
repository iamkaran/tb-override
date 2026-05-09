// Package edit contains method to add or remove variable overrides from custom CSS files
package edit

import (
	"fmt"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/fs"
)

func AppendVariable(cfg *config.Config, theme string, variableName, value string) error {
	cssFilePath := cfg.TBOverride.Dirs.RootDirectory +
		"/" +
		cfg.TBOverride.Dirs.ThemesDirectory +
		"/" +
		theme +
		"/" +
		cfg.TBOverride.Files.CSSFilename

	data, err := fs.GetFileContents(cssFilePath)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
