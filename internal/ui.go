package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// PrintHelp displays usage information.
func PrintHelp() {
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
	fmt.Println("\nEnvironment Aliases:")
	fmt.Println("  prod, p, live     → production")
	fmt.Println("  stg, stage        → staging")
	fmt.Println("  d, local          → dev")
}

// ListConfig displays all configured URLs for the current project.
func ListConfig() {
	cfg := LoadConfig()
	if len(cfg) == 0 {
		fmt.Println("No URLs configured.")
		return
	}
	fmt.Println("Configured URLs:")
	for env, url := range cfg {
		fmt.Printf("%-12s: %s\n", env, url)
	}
}

// RunInteractiveSetup launches the guided CLI experience.
func RunInteractiveSetup() {
	fmt.Println("🌟 gourl Interactive Setup")
	fmt.Println("This guide will help you configure URLs for this project.")

	// 1. Detect project type
	projectType, defaultPort := DetectProject()
	if projectType != "" {
		fmt.Printf("🔍 Detected %s project. Suggested dev port: %s\n", projectType, defaultPort)
	}

	scanner := bufio.NewScanner(os.Stdin)
	envs := []string{"dev", "staging", "production"}
	cfg := LoadConfig()

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
			cfg[env] = input
		}
	}

	// 2. Save configuration
	os.MkdirAll(".cache", 0755)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(ConfigFileName, data, 0644)

	fmt.Println("\n✅ Configuration saved to .cache/gourls.json")
	ListConfig()
	CheckAndSuggestGitignore()
}

// CheckAndSuggestGitignore helps ensure the .cache/ is not committed.
func CheckAndSuggestGitignore() {
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
