// Package apply contains ApplyTheme method that writes the theme into state.json to make it active
package apply

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

func ApplyTheme(log *slog.Logger, cfg *config.Config, themeName string) error {
	themesDir := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ThemesDirectory,
	)

	stateFile := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Files.StateFile,
	)

	listOfThemes, err := fs.ListDirs(themesDir)
	if err != nil {
		return err
	}
	isValid := false
	for _, theme := range listOfThemes {
		if theme == themeName {
			isValid = true
		}
	}
	if !isValid {
		return core.ErrInvalidTheme
	}

	data := core.JSONState{
		ActiveTheme: themeName,
	}

	fileData, _ := json.MarshalIndent(data, "", "    ")
	err = fs.WriteToFile(stateFile, fileData)
	if err != nil {
		return err
	}

	fmt.Printf("Theme %s is set to active\n", themeName)

	return nil
}
