package fs

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/iamkaran/tb-override/internal/core"
)

func CreateFile(log *slog.Logger, path string) error {
	if os.Geteuid() != 0 && requireRootPrivilages(path, []string{"/var", "/usr", "/etc"}) {
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
