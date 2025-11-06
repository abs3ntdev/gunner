// Package gunner provides a simple configuration loader for Go applications.
// It supports loading configuration from YAML, JSON, and .env files with
// environment variable overrides.
package gunner

import (
	"os"
	"path/filepath"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"

	"github.com/abs3ntdev/gunner/src/yaml"
)

// LoadApp loads application configuration from multiple sources.
// It searches for configuration files in the user's config directory
// (typically ~/.config/{appName}/) and supports YAML, JSON, and .env formats.
// Environment variables with the appName prefix can override file-based config.
//
// Parameters:
//   - i: pointer to a struct that will be populated with configuration values
//   - appName: name of the application, used for config directory and env var prefix
//
// Returns an error if configuration loading fails.
func LoadApp(i any, appName string) error {
	yamlDecoder := yaml.New()
	dotenvDecoder := aconfigdotenv.New()

	configDir, err := os.UserConfigDir()
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir = filepath.Join(homeDir, ".config")
	}

	filePath := filepath.Join(configDir, appName)

	loader := aconfig.LoaderFor(i, aconfig.Config{
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		AllowUnknownFlags:  true,
		SkipFlags:          true,
		DontGenerateTags:   true,
		MergeFiles:         true,
		EnvPrefix:          appName,
		FlagPrefix:         appName,
		Files: []string{
			filepath.Join(filePath, appName+".yml"),
			filepath.Join(filePath, appName+".yaml"),
			filepath.Join(filePath, appName+".json"),
			filepath.Join(filePath, ".env"),
		},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": yamlDecoder,
			".yml":  yamlDecoder,
			".json": yamlDecoder,
			".env":  dotenvDecoder,
		},
	})
	if err := loader.Load(); err != nil {
		return err
	}
	return nil
}
