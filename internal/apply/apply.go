// Package apply contains ApplyTheme method that writes the theme into state.json to make it active
package apply

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

func ApplyTheme(cfg *config.Config, themeName string) error {
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

	if !slices.Contains(listOfThemes, themeName) {
		return core.ErrInvalidTheme
	}

	data := core.JSONState{
		ActiveTheme: themeName,
	}

	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = ActivateCSSFile(cfg, themeName)
	if err != nil {
		return err
	}

	err = fs.WriteToFile(stateFile, fileData)
	if err != nil {
		return err
	}

	fmt.Printf("Theme %s is set to active\n", themeName)

	return nil
}

func ActivateCSSFile(cfg *config.Config, activeTheme string) error {
	themeCSSPath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ThemesDirectory,
		activeTheme,
		cfg.TBOverride.Files.CSSFilename,
	)

	activeCSSPath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ActiveDirectory,
		cfg.TBOverride.Files.CSSFilename,
	)

	relativeCSSPath, err := filepath.Rel(
		filepath.Dir(activeCSSPath),
		themeCSSPath,
	)

	if err != nil {
		return err
	}

	err = os.Remove(activeCSSPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = os.Symlink(relativeCSSPath, activeCSSPath)
	if err != nil {
		return err
	}

	return nil
}
