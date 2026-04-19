package internal

import (
	"os"
)

// Info represents detected project traits.
type Info struct {
	Name string
	Port string
}

// DetectProject scans the current directory for well-known framework markers.
func DetectProject() (string, string) {
	markers := map[string]struct {
		name string
		port string
	}{
		"go.mod":           {"Go", "8080"},
		"package.json":     {"Node.js", "3000"},
		"yarn.lock":        {"Node.js (Yarn)", "3000"},
		"pnpm-lock.yaml":   {"Node.js (pnpm)", "3000"},
		"bun.lockb":        {"Node.js (Bun)", "3000"},
		"Gemfile":          {"Ruby/Rails", "3000"},
		"requirements.txt": {"Python", "8000"},
		"Cargo.toml":       {"Rust", "8080"},
	}

	for marker, info := range markers {
		if _, err := os.Stat(marker); err == nil {
			return info.name, info.port
		}
	}
	return "", ""
}
