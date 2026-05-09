// Package config is responsible for Initializing, Parsing, Loading of various types of config variables throught the package Viper
package config

import (
	"context"
	"errors"
	"github.com/iamkaran/tb-override/internal/core"
	"github.com/iamkaran/tb-override/internal/detect"
	"github.com/iamkaran/tb-override/internal/logger"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Thingsboard ThingsboardConfig `mapstructure:"thingsboard"`
	Logger      LoggerConfig      `mapstructure:"logger"`
	TBOverride  TBOverrideConfig  `mapstructure:"tb-override"`
	Output      OutputConfig      `mapstructure:"output"`
}

type ThingsboardConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type TBOverrideConfig struct {
	Files TBOverrideFilesConfig `mapstructure:"files"`
	Dirs  TBOverrideDirsConfig  `mapstructure:"dirs"`
	Misc  TBOverrideMiscConfig  `mapstructure:"misc"`
}

type TBOverrideFilesConfig struct {
	NginxConfig       string `mapstructure:"nginx_config"`
	StateFile         string `mapstructure:"state_file"`
	CSSFilename       string `mapstructure:"css_filename"`
	VariablesFilename string `mapstructure:"variables_filename"`
}

type TBOverrideDirsConfig struct {
	RootDirectory   string `mapstructure:"root_directory"`
	ThemesDirectory string `mapstructure:"themes_directory"`
	ActiveDirectory string `mapstructure:"active_directory"`
}

type TBOverrideMiscConfig struct {
	SkipProxyCheck bool `mapstructure:"skip_proxy_check"`
}

type LoggerConfig struct {
	Format string `mapstructure:"format"`
	Level  string `mapstructure:"level"`
}

type OutputConfig struct {
	Verbose bool `mapstructure:"verbose"`
}

// InitializeConfig Initializes Config the config from various means like Env, Config file, Flags
// Loads the config into a reusable struct
func InitializeConfig(cfgFile string, cmd *cobra.Command) error {
	viper.SetEnvPrefix("TBOVERRIDE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Checks the following paths for config
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.tb-override")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	log := logger.New(cfg.Logger.Level, cfg.Logger.Format)
	ctx = context.WithValue(ctx, core.LoggerKey, log)

	platform, err := detect.PlatformInfo()
	if err != nil && !cfg.TBOverride.Misc.SkipProxyCheck {
		log.Error("NGINX Not found $PATH")
	}

	ctx = context.WithValue(ctx, core.ConfigKey, cfg)

	ctx = context.WithValue(ctx, core.PlatformKey, platform)

	log.Debug("config object", "config", cfg)
	cmd.SetContext(ctx)
	return nil
}

func FromContext(ctx context.Context) *Config {
	c, ok := ctx.Value(core.ConfigKey).(*Config)
	if !ok {
		log.Fatalf("Config could not be fetched: %v\n", ok)
		return nil
	}

	return c
}
