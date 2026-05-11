// Package fs holds helper functions for creating directories, files and safely editing config files
package fs

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
)

func CreateDir(log *slog.Logger, cfg *config.Config, path string) error {
	err := unix.Access(cfg.TBOverride.Dirs.RootDirectory, unix.W_OK)
	if err != nil {
		return core.ErrNoRootPrivilages
	}

	cleanPath := filepath.Clean(path)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Error("home directory error", "error", err)
	}

	if strings.HasPrefix(cleanPath, "~") {
		cleanPath = strings.Replace(cleanPath, "~", homeDir, 1)
	}

	if _, err := os.Stat(cleanPath); !errors.Is(err, os.ErrNotExist) {
		return core.ErrAlreadyExists
	}

	err = os.MkdirAll(cleanPath, 0755)

	if err != nil {
		return err
	}

	return nil
}

func ListDirs(path string) ([]string, error) {
	cleanPath := filepath.Clean(path)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(cleanPath, "~") {
		cleanPath = strings.Replace(cleanPath, "~", homeDir, 1)
	}
	entries, err := os.ReadDir(cleanPath)
	directories := []string{}
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.IsDir() {
			directories = append(directories, e.Name())
		}
	}

	return directories, nil
}
