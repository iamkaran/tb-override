// Package apply contains ApplyTheme method that writes the theme into state.json to make it active
package apply

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

type JSONState struct {
	ActiveTheme string `json:"active_theme"`
}

func ApplyTheme(log *slog.Logger, cfg *config.Config, themeName string) error {
	listOfThemes, err := fs.ListDirs(cfg.TBOverride.ThemesDirectory)
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

	data := JSONState{
		ActiveTheme: themeName,
	}

	fileData, _ := json.MarshalIndent(data, "", "    ")
	err = fs.WriteToFile(cfg.TBOverride.StateFile, fileData)
	if err != nil {
		return err
	}

	fmt.Printf("Theme %s is set to active\n", themeName)

	return nil
}
