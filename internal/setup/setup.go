// Package setup contains the logic for setting up tb-override for the first time
package setup

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

func Setup(ctx context.Context, log *slog.Logger, cfg *config.Config) error {
	err := setupFilesAndDirs(cfg, log)
	if err != nil {
		return err
	}

	err = setupStateFile(cfg)
	if err != nil {
		return err
	}

	err = setupVariables(cfg)
	if err != nil {
		return err
	}

	err = setupCSSFile(cfg)
	if err != nil {
		return err
	}

	return nil
}

func setupCSSFile(cfg *config.Config) error {
	cssPath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Files.RulesFilename,
	)

	if _, err := os.Stat(cssPath); errors.Is(err, os.ErrNotExist) {
		src, err := os.Open(cfg.TBOverride.Files.ExampleRulesFilename)
		if err != nil {
			return err
		}
		defer func() {
			_ = src.Close()
		}()

		dst, err := os.Create(cssPath)
		if err != nil {
			return err
		}

		defer func() {
			_ = src.Close()
		}()

		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupStateFile(cfg *config.Config) error {
	stateFile := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Files.StateFile,
	)

	if _, err := os.Stat(stateFile); errors.Is(err, os.ErrNotExist) {

		data := core.JSONState{
			ActiveTheme: "",
		}

		fileData, _ := json.MarshalIndent(data, "", "    ")
		err = fs.WriteToFile(stateFile, fileData)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupVariables(cfg *config.Config) error {
	variablesPath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Files.VariablesFilename,
	)

	if _, err := os.Stat(variablesPath); errors.Is(err, os.ErrNotExist) {
		src, err := os.Open(cfg.TBOverride.Files.ExampleVariablesFilename)
		if err != nil {
			return err
		}
		defer func() {
			_ = src.Close()
		}()

		dst, err := os.Create(variablesPath)
		if err != nil {
			return err
		}

		defer func() {
			_ = src.Close()
		}()

		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupFilesAndDirs(cfg *config.Config, log *slog.Logger) error {
	themesDir := filepath.Join(cfg.TBOverride.Dirs.RootDirectory, cfg.TBOverride.Dirs.ThemesDirectory)
	activeDir := filepath.Join(cfg.TBOverride.Dirs.RootDirectory, cfg.TBOverride.Dirs.ActiveDirectory)

	directories := []string{
		themesDir,
		activeDir,
	}

	nginxConfig := filepath.Join(cfg.TBOverride.Dirs.RootDirectory, cfg.TBOverride.Files.NginxConfig)
	stateFile := filepath.Join(cfg.TBOverride.Dirs.RootDirectory, cfg.TBOverride.Files.StateFile)

	files := []string{
		nginxConfig,
		stateFile,
	}

	err := validateAndCreate(log, cfg, "directory", directories)
	if err != nil {
		return err
	}

	err = validateAndCreate(log, cfg, "file", files)
	if err != nil {
		return err
	}

	return nil

}

func validateAndCreate(log *slog.Logger, cfg *config.Config, actionType string, objects []string) error {
	for _, path := range objects {
		cleanPath := filepath.Clean(path)

		var err error
		switch actionType {
		case "file":
			err = fs.CreateFile(log, cfg, cleanPath)
		case "directory":
			err = fs.CreateDir(log, cfg, cleanPath)
		}

		if errors.Is(err, core.ErrNoRootPrivilages) {
			log.Error(core.ErrNoRootPrivilages.Error(), actionType, cleanPath)
			return core.ErrNoRootPrivilages
		} else if errors.Is(err, core.ErrAlreadyExists) {
			log.Debug("skipping creation, already exists", actionType, cleanPath)
		} else if err != nil {
			log.Error("failed to create", actionType, cleanPath, "error", err)
			return err
		} else {
			log.Debug("successfully created", actionType, cleanPath)
		}
	}
	return nil
}
