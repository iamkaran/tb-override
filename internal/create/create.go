// Package create has method to create a theme with the starter custom.css
package create

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/fs"
)

func CreateTheme(log *slog.Logger, cfg *config.Config, themeName string) error {
	themesDirectory := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ThemesDirectory,
	)

	customThemeDirectory := filepath.Join(
		themesDirectory,
		themeName,
	)

	cssFilePath := filepath.Join(
		customThemeDirectory,
		cfg.TBOverride.Files.CSSFilename,
	)

	if err := fs.CreateDir(log, cfg, customThemeDirectory); err != nil {
		return err
	}

	if err := fs.CreateFile(log, cfg, cssFilePath); err != nil {
		return err
	}

	defaultCSS := []byte(`:root {

}`)

	err := fs.WriteToFile(cssFilePath, defaultCSS)
	if err != nil {
		return err
	}

	fmt.Printf("Theme %s created at %s\n", themeName, customThemeDirectory)

	return nil
}
