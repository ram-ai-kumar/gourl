package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const configPath = ".cache/gourls.json"

var Version = "dev"

type Config map[string]string

func main() {
	args := os.Args[1:]

	// 1. Default 'gourl' -> Show help
	if len(args) == 0 {
		printHelp()
		return
	}

	command := args[0]

	switch command {
	case "set":
		if len(args) < 3 {
			fmt.Println("Usage: gourl set <env> <url>")
			return
		}
		saveConfig(args[1], args[2])
	case "list":
		listConfig()
	case "version", "--version", "-v":
		fmt.Printf("gourl version %s\n", Version)
	case "help", "--help", "-h":
		printHelp()
	default:
		// 2. 'gourl <env>' -> Open specific env
		openUrl(command)
	}
}

func normalizeEnv(env string) string {
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

func loadConfig() Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return make(Config)
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

func saveConfig(env, url string) {
	os.MkdirAll(".cache", 0755)
	cfg := loadConfig()
	env = normalizeEnv(env)
	cfg[env] = url

	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configPath, data, 0644)
	fmt.Printf("✅ Saved %s -> %s\n", env, url)
}

func openUrl(env string) {
	env = normalizeEnv(env)
	cfg := loadConfig()
	url, ok := cfg[env]

	if !ok {
		fmt.Printf("❌ No URL found for '%s'. Run: gourl set %s <url>\n", env, env)

		// Check if config file exists and suggest gitignore
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Println("💡 No URLs configured for this project.")
			checkAndSuggestGitignore()
		}
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		// Check for WSL
		if _, err := exec.LookPath("wslview"); err == nil {
			cmd = exec.Command("wslview", url)
		} else {
			cmd = exec.Command("xdg-open", url)
		}
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("❌ Error opening URL: %v\n", err)
	} else {
		fmt.Printf("🚀 Opening %s\n", url)
	}
}

func listConfig() {
	cfg := loadConfig()
	if len(cfg) == 0 {
		fmt.Println("No URLs configured.")
		return
	}
	fmt.Println("Configured URLs:")
	for env, url := range cfg {
		fmt.Printf("%-12s: %s\n", env, url)
	}
}

func printHelp() {
	fmt.Println("gourl - Project URL Manager")
	fmt.Println("\nUsage:")
	fmt.Println("  gourl              Show this help message")
	fmt.Println("  gourl <env>        Opens specific URL (e.g., prod, staging, dev)")
	fmt.Println("  gourl set <env> <url>   Save a URL")
	fmt.Println("  gourl list         List all saved URLs")
	fmt.Println("  gourl version      Show version information")
	fmt.Println("  gourl help         Show this help message")
	fmt.Println("\nEnvironment Aliases:")
	fmt.Println("  prod, p, live     → production")
	fmt.Println("  stg, stage        → staging")
	fmt.Println("  d, local          → dev")
}

func checkAndSuggestGitignore() {
	// Check if .cache is gitignored
	gitignorePath := ".gitignore"
	if _, err := os.Stat(gitignorePath); err == nil {
		content, err := os.ReadFile(gitignorePath)
		if err == nil && strings.Contains(string(content), ".cache/") {
			return // Already gitignored
		}
	}

	fmt.Println("💡 To ignore .cache/ in git, run:")
	fmt.Println("   echo '.cache/' >> .gitignore")
}
