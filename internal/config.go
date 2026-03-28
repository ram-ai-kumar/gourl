package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const LocalConfigPath = ".cache/gourls.json"

// Config represents a mapping of environments to URLs.
type Config struct {
	Envs     map[string]string `json:"envs,omitempty"`     // Environment mappings
	Defaults map[string]string `json:"defaults,omitempty"` // Project-type defaults (global only)
}

// GetGlobalConfigPath returns the absolute path to ~/.cache/gourls.json.
func GetGlobalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache", "gourls.json"), nil
}

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

// LoadConfig loads the configuration from a specific path.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Config{Envs: make(map[string]string), Defaults: make(map[string]string)}, err
	}
	var cfg Config
	// Support both legacy (flat map) and new (structured) formats
	var flatMap map[string]string
	if err := json.Unmarshal(data, &flatMap); err == nil {
		// If it's a flat map, migrate it to Envs
		return &Config{Envs: flatMap, Defaults: make(map[string]string)}, nil
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return &Config{Envs: make(map[string]string), Defaults: make(map[string]string)}, err
	}
	if cfg.Envs == nil {
		cfg.Envs = make(map[string]string)
	}
	if cfg.Defaults == nil {
		cfg.Defaults = make(map[string]string)
	}
	return &cfg, nil
}

// SaveConfig saves the configuration to a specific path.
func SaveConfig(path string, cfg *Config) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GetConfigValue resolves a URL checking local, then global, then project defaults.
func GetConfigValue(env string) (string, bool) {
	env = NormalizeEnv(env)

	// 1. Check local config
	localCfg, _ := LoadConfig(LocalConfigPath)
	if url, ok := localCfg.Envs[env]; ok {
		return url, true
	}

	// 2. Check global config
	globalPath, err := GetGlobalConfigPath()
	if err == nil {
		globalCfg, _ := LoadConfig(globalPath)
		if url, ok := globalCfg.Envs[env]; ok {
			return url, true
		}

		// 3. Check for project defaults in global (only for 'dev')
		if env == "dev" {
			projectType, _ := DetectProject()
			if url, ok := globalCfg.Defaults[projectType]; ok {
				return url, true
			}
		}
	}

	return "", false
}

// UnsetConfig removes a mapping from local or global config.
func UnsetConfig(env string, isGlobal bool) error {
	env = NormalizeEnv(env)
	path := LocalConfigPath
	if isGlobal {
		var err error
		path, err = GetGlobalConfigPath()
		if err != nil {
			return err
		}
	}

	cfg, _ := LoadConfig(path)
	delete(cfg.Envs, env)
	return SaveConfig(path, cfg)
}

// PreseedDefaults ensures standard defaults are present in the global config.
func PreseedDefaults(cfg *Config) {
	standardDefaults := map[string]string{
		"Go":             "http://localhost:8080",
		"Node.js":        "http://localhost:3000",
		"Node.js (Yarn)": "http://localhost:3000",
		"Node.js (pnpm)": "http://localhost:3000",
		"Node.js (Bun)":  "http://localhost:3000",
		"Ruby/Rails":     "http://localhost:3000",
		"Python":         "http://localhost:8000",
		"Rust":           "http://localhost:8080",
	}

	if cfg.Defaults == nil {
		cfg.Defaults = make(map[string]string)
	}

	for k, v := range standardDefaults {
		if _, ok := cfg.Defaults[k]; !ok {
			cfg.Defaults[k] = v
		}
	}
}
