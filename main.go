// Package main provides the gourl CLI, a secure local-first URL manager for developers.
// It allows mapping environment names (prod, staging, dev) to project-specific URLs
// and opening them in the default browser with zero trust principles.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
			os.Exit(1)
		}
		saveConfig(args[1], args[2])
	case "list":
		listConfig()
	case "version", "--version", "-v":
		fmt.Printf("gourl version %s\n", Version)
	case "help", "--help", "-h":
		printHelp()
	case "-i", "--interactive":
		runInteractiveSetup()
	case "--purge":
		force := false
		for _, arg := range args {
			if arg == "--force" {
				force = true
			}
		}
		purge(force)
	default:
		// Check if any argument is --purge (in case it's not the first)
		isPurge := false
		for _, arg := range args {
			if arg == "--purge" {
				isPurge = true
				break
			}
		}
		if isPurge {
			force := false
			for _, arg := range args {
				if arg == "--force" {
					force = true
				}
			}
			purge(force)
			return
		}
		// 2. 'gourl <env>' -> Open specific env
		openUrl(command)
	}
}

// purge removes the gourl binary from the system.
// If force is false, it prompts the user for confirmation.
// project-specific .cache/ directories are explicitly preserved.
func purge(force bool) {
	// 1. Locate the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("❌ Error: Could not locate gourl executable: %v\n", err)
		os.Exit(1)
	}

	// 2. Resolve symlinks
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		realPath = exePath // Fallback to raw path if symlink resolution fails
	}

	// 3. Confirmation
	if !force {
		fmt.Printf("⚠️  Are you sure you want to uninstall gourl? (this will remove the binary at %s) [y/N]: ", realPath)
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" && response != "yes" {
			fmt.Println("❌ Uninstallation cancelled.")
			return
		}
	}

	// 4. Remove the binary
	err = os.Remove(realPath)
	if err != nil {
		fmt.Printf("❌ Error: Failed to remove binary: %v\n", err)
		if runtime.GOOS == "windows" {
			fmt.Println("💡 Note: On Windows, you may need to close other instances of gourl before uninstalled.")
		}
		os.Exit(1)
	}

	fmt.Println("✅ gourl has been successfully uninstalled.")
	fmt.Println("💡 Note: Your project-specific .cache/ folders were preserved.")
}

// runInteractiveSetup launches a guided CLI experience to configure project URLs.
// It detects the framework (Go, Node, etc.) and suggests default development ports.
func runInteractiveSetup() {
	fmt.Println("🌟 gourl Interactive Setup")
	fmt.Println("This guide will help you configure URLs for this project.")

	// 1. Detect project type
	projectType, defaultPort := detectProject()
	if projectType != "" {
		fmt.Printf("🔍 Detected %s project. Suggested dev port: %s\n", projectType, defaultPort)
	}

	scanner := bufio.NewScanner(os.Stdin)
	envs := []string{"dev", "staging", "production"}
	config := loadConfig()

	for _, env := range envs {
		defaultUrl := ""
		if env == "dev" && defaultPort != "" {
			defaultUrl = fmt.Sprintf("http://localhost:%s", defaultPort)
		}

		prompt := fmt.Sprintf("Enter URL for %s", env)
		if defaultUrl != "" {
			prompt += fmt.Sprintf(" [%s]", defaultUrl)
		}
		fmt.Printf("%s: ", prompt)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" && defaultUrl != "" {
			input = defaultUrl
		}

		if input != "" {
			config[env] = input
		}
	}

	// 2. Save configuration
	os.MkdirAll(".cache", 0755)
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configPath, data, 0644)

	fmt.Println("\n✅ Configuration saved to .cache/gourls.json")
	listConfig()
	checkAndSuggestGitignore()
}

// detectProject scans the current directory for well-known framework markers.
// Returns the framework name and its standard development port.
func detectProject() (string, string) {
	markers := map[string]struct {
		name string
		port string
	}{
		"go.mod":           {"Go", "8080"},
		"package.json":     {"Node.js", "3000"},
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
		os.Exit(1)
	}

	// Check if we're in test mode and stub the opening
	if os.Getenv("GOURL_TEST_MODE") == "1" {
		fmt.Printf("🧪 TEST MODE: Would open %s (stubbed)\n", url)
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
	fmt.Println("  gourl -i, --interactive   Guided setup for project URLs")
	fmt.Println("  gourl --purge      Uninstall gourl (removes the binary)")
	fmt.Println("  gourl --purge --force  Uninstall without confirmation")
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
