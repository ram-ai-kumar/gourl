package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const ConfigFileName = ".cache/gourls.json"

type Config map[string]string

// NormalizeEnv maps environment shorthand to standard keys.
func NormalizeEnv(env string) string {
	switch strings.ToLower(env) {
	case "prod", "p", "live":
		return "production"
	case "stg", "stage":
		return "staging"
	case "d", "local":
		return "dev"
	default:
		return env
	}
}

// LoadConfig loads the configuration from the .cache directory.
func LoadConfig() Config {
	data, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return make(Config)
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

// SaveConfig saves the provided URL to the environment mapping.
func SaveConfig(env, url string) error {
	os.MkdirAll(filepath.Dir(ConfigFileName), 0755)
	cfg := LoadConfig()
	env = NormalizeEnv(env)
	cfg[env] = url

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigFileName, data, 0644)
}
