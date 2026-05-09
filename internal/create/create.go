package create

import (
	"fmt"
	"log/slog"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/fs"
)

func CreateTheme(log *slog.Logger, cfg *config.Config, themeName string) error {
	themesDirectory := cfg.TBOverride.ThemesDirectory
	customThemeDirectory := fmt.Sprintf("%s/%s", themesDirectory, themeName)
	cssFilePath := fmt.Sprintf("%s/%s", customThemeDirectory, cfg.TBOverride.CSSFilename)

	if err := fs.CreateDir(log, customThemeDirectory); err != nil {
		return err
	}
	if err := fs.CreateFile(log, cssFilePath); err != nil {
		return err
	}

	fmt.Printf("Theme %s created at %s\n", themeName, customThemeDirectory)

	return nil
}
