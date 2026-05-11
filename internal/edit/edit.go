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

func SetVariable(cfg *config.Config, property core.CSSProperty) error {
	cssContents, cssFilePath, err := GetCSSContents(cfg)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(
		fmt.Sprintf(`%s:\s*.*?;`, regexp.QuoteMeta(property.Name)),
	)

	if re.MatchString(cssContents) {
		return EditVariable(cfg, property)
	}

	propertyStr := fmt.Sprintf("  %s: %s;\n", property.Name, property.Value)
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

func EditVariable(cfg *config.Config, property core.CSSProperty) error {
	cssContents, cssFilePath, err := GetCSSContents(cfg)
	if err != nil {
		return err
	}

	propertyStr := fmt.Sprintf("  %s: %s;", property.Name, property.Value)

	re := regexp.MustCompile(fmt.Sprintf(`%s:\s*.*?;`, regexp.QuoteMeta(property.Name)))

	if !re.MatchString(cssContents) {
		return core.ErrCSSPropNotExist
	}

	updated := re.ReplaceAllString(
		cssContents,
		propertyStr,
	)

	err = fs.WriteToFile(cssFilePath, []byte(updated))
	if err != nil {
		return err
	}

	return nil
}

func DeleteVariable(cfg *config.Config, property core.CSSProperty) error {
	cssContents, cssFilePath, err := GetCSSContents(cfg)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(fmt.Sprintf(`\s*%s:\s*.*?;\n?`, regexp.QuoteMeta(property.Name)))
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

func GetCSSContents(cfg *config.Config) (string, string, error) {
	cssFileName, err := fs.GetActiveCSS(cfg)
	if err != nil {
		return "", "", err
	}

	cssFilePath := filepath.Join(
		cfg.TBOverride.Dirs.RootDirectory,
		cfg.TBOverride.Dirs.ThemesDirectory,
		cssFileName,
	)

	cssContents, err := fs.GetFileContents(cssFilePath)
	if err != nil {
		return "", "", err
	}

	return cssContents, cssFilePath, nil
}
