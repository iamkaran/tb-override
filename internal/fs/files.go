package fs

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"golang.org/x/sys/unix"
)

func CreateFile(log *slog.Logger, cfg *config.Config, path string) error {
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

	_, err = os.Create(cleanPath)

	if err != nil {
		return err
	}

	return nil
}

func GetFileContents(path string) (string, error) {
	cleanPath := filepath.Clean(path)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(cleanPath, "~") {
		cleanPath = strings.Replace(cleanPath, "~", homeDir, 1)
	}
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func WriteToFile(path string, data []byte) error {
	cleanPath := filepath.Clean(path)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if strings.HasPrefix(cleanPath, "~") {
		cleanPath = strings.Replace(cleanPath, "~", homeDir, 1)
	}

	err = os.WriteFile(cleanPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetActiveCSS(cfg *config.Config) (string, error) {
	stateFile, err := os.ReadFile(cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.StateFile)
	if err != nil {
		return "", err
	}

	state := core.JSONState{}
	err = json.Unmarshal(stateFile, &state)
	if err != nil {
		return "", err
	}

	if state.ActiveTheme == "" {
		return "", core.ErrNoActiveTheme
	}

	return state.ActiveTheme, nil
}
