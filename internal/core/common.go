package core

type contextKey string

const (
	ConfigKey   contextKey = "config"
	PlatformKey contextKey = "platform"
	LoggerKey   contextKey = "logger" // Is imported in logger.go
)

type JSONState struct {
	ActiveTheme string `json:"active_theme"`
}
