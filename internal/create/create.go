package create

import (
	"fmt"
	"log/slog"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/fs"
)

func CreateTheme(log *slog.Logger, cfg *config.Config, themeName string) error {
	themesDirectory := cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Dirs.ThemesDirectory
	customThemeDirectory := themesDirectory + "/" + themeName
	cssFilePath := customThemeDirectory + "/" + cfg.TBOverride.Files.CSSFilename

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
