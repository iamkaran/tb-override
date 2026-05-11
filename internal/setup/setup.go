// Package install contains the logic for setting up tb-override for the first time
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
	// tb-override nginx directory
	directories := []string{
		cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Dirs.ThemesDirectory,
		cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Dirs.ActiveDirectory,
	}

	files := []string{
		cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.NginxConfig,
		cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.StateFile,
	}

	err := validateAndCreate(log, cfg, "directory", directories)
	if err != nil {
		return err
	}

	err = validateAndCreate(log, cfg, "file", files)
	if err != nil {
		return err
	}

	variablesPath := "example_variables.json"

	if _, err := os.Stat(cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.VariablesFilename); errors.Is(err, os.ErrNotExist) {
		src, err := os.Open(variablesPath)
		if err != nil {
			return err
		}
		defer func() {
			_ = src.Close()
		}()

		dst, err := os.Create(cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.VariablesFilename)
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
	if _, err := os.Stat(cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.StateFile); errors.Is(err, os.ErrNotExist) {

		data := core.JSONState{
			ActiveTheme: "",
		}

		fileData, _ := json.MarshalIndent(data, "", "    ")
		err = fs.WriteToFile(cfg.TBOverride.Dirs.RootDirectory+"/"+cfg.TBOverride.Files.StateFile, fileData)
		if err != nil {
			return err
		}
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
