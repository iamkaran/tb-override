// Package install contains the logic for setting up tb-override for the first time
package setup

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

func Setup(ctx context.Context, log *slog.Logger, cfg *config.Config) error {
	// tb-override nginx directory
	directories := []string{
		cfg.TBOverride.ThemesDirectory,
		cfg.TBOverride.ActiveDirectory,
	}

	files := []string{
		cfg.TBOverride.NginxConfig,
		cfg.TBOverride.StateFile,
	}

	err := validateAndCreate(log, "directory", directories)
	if err != nil {
		return err
	}

	err = validateAndCreate(log, "file", files)
	if err != nil {
		return err
	}

	fmt.Println("Required directories and files have been created")

	return nil
}

func validateAndCreate(log *slog.Logger, actionType string, objects []string) error {
	for _, path := range objects {
		cleanPath := filepath.Clean(path)

		var err error
		switch actionType {
		case "file":
			err = fs.CreateFile(log, cleanPath)
		case "directory":
			err = fs.CreateDir(log, cleanPath)
		}

		if errors.Is(err, core.ErrNoRootPrivilages) {
			log.Error(core.ErrNoRootPrivilages.Error(), actionType, cleanPath)
			return err
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
