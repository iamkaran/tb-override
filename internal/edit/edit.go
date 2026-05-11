// Package edit contains method to add or remove variable overrides from custom CSS files
package edit

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/fs"
)

func SetVariable(cfg *config.Config, themeName string, property core.CSSProperty) error {
	cssContents, cssFilePath, err := GetCSSContents(cfg, themeName)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(
		fmt.Sprintf(`[ \t]*%s:\s*.*?;`, regexp.QuoteMeta(property.Name)),
	)

	if re.MatchString(cssContents) {
		propertyStr := fmt.Sprintf("  %s: %s;", property.Name, property.Value)

		updated := re.ReplaceAllString(
			cssContents,
			propertyStr,
		)

		return fs.WriteToFile(cssFilePath, []byte(updated))
	}

	propertyStr := fmt.Sprintf("\n  %s: %s;\n", property.Name, property.Value)
	idx := strings.LastIndex(cssContents, "}")

	if idx == -1 {
		return core.ErrInvalidCSS
	}

	updated := cssContents[:idx] +
		propertyStr +
		cssContents[idx:]

	err = fs.WriteToFile(cssFilePath, []byte(updated))
	if err != nil {
		return err
	}

	return nil
}

func DeleteVariable(cfg *config.Config, themeName string, variableName string) error {
	cssContents, cssFilePath, err := GetCSSContents(cfg, themeName)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(fmt.Sprintf(`\s*%s:\s*.*?;\n?`, regexp.QuoteMeta(variableName)))
	updated := re.ReplaceAllString(
		cssContents,
		"",
	)

	err = fs.WriteToFile(cssFilePath, []byte(updated))
	if err != nil {
		return err
	}

	return nil
}

func GetCSSContents(cfg *config.Config, themeName string) (string, string, error) {
	cssFilePath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ThemesDirectory,
		themeName,
		cfg.TBOverride.Files.CSSFilename,
	)

	cssContents, err := fs.GetFileContents(cssFilePath)
	if err != nil {
		return "", "", err
	}

	return cssContents, cssFilePath, nil
}
