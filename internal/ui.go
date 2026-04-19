package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PrintHelp displays usage information.
func PrintHelp() {
	fmt.Println("gourl - Project URL Manager")
	fmt.Println("\nUsage:")
	fmt.Println("  gourl                    Show this help message")
	fmt.Println("  gourl <env>              Opens specific URL (e.g., prod, staging, dev)")
	fmt.Println("  gourl set <env> <url>    Save a URL to local project config")
	fmt.Println("  gourl set -g <env> <url> Save a URL to global config (~/.cache/gourls.json)")
	fmt.Println("  gourl set -f <env> <url> Save a URL to favourites (~/.cache/gourl-favourites.json)")
	fmt.Println("  gourl unset <env>        Remove a URL from local project config")
	fmt.Println("  gourl unset -g <env>     Remove a URL from global config")
	fmt.Println("  gourl unset -f <env>     Remove a URL from favourites")
	fmt.Println("  gourl list               List all URLs (merged local, global, and favourites)")
	fmt.Println("  gourl version            Show version information")
	fmt.Println("  gourl help               Show this help message")
	fmt.Println("  gourl -i, --interactive  Guided setup for project URLs")
	fmt.Println("  gourl --purge            Uninstall gourl (removes the binary)")
	fmt.Println("\nEnvironment Aliases:")
	fmt.Println("  prod, p, live     → production")
	fmt.Println("  stg, stage        → staging")
	fmt.Println("  d, local          → dev")
}

// ListConfig displays all configured URLs for the current project context.
func ListConfig() {
	// Merged view
	localCfg, _ := LoadConfig(LocalConfigPath)
	globalPath, _ := GetGlobalConfigPath()
	globalCfg, _ := LoadConfig(globalPath)
	favPath, _ := GetFavouritesConfigPath()
	favCfg, _ := LoadConfig(favPath)

	// Combine them for listing
	merged := make(map[string]struct {
		url    string
		source string
	})

	// Global first
	for env, url := range globalCfg.Envs {
		merged[env] = struct {
			url    string
			source string
		}{url: url, source: "global"}
	}

	// Local overrides
	for env, url := range localCfg.Envs {
		merged[env] = struct {
			url    string
			source string
		}{url: url, source: "local"}
	}

	// Favourites
	for env, url := range favCfg.Favourites {
		merged[env] = struct {
			url    string
			source string
		}{url: url, source: "favourite"}
	}

	// If dev is missing, check project default
	if _, ok := merged["dev"]; !ok {
		projectType, _ := DetectProject()
		if url, ok := globalCfg.Defaults[projectType]; ok {
			merged["dev"] = struct {
				url    string
				source string
			}{url: url, source: "default"}
		}
	}

	if len(merged) == 0 {
		fmt.Println("No URLs configured.")
		return
	}

	fmt.Println("Configured URLs:")
	for env, info := range merged {
		fmt.Printf("%-12s: %-30s (%s)\n", env, info.url, info.source)
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
	cfg, _ := LoadConfig(LocalConfigPath)

	for _, env := range envs {
		// Use existing value or project default as suggestion
		existingUrl, ok := cfg.Envs[env]
		suggestion := ""
		if ok {
			suggestion = existingUrl
		} else if env == "dev" && defaultPort != "" {
			suggestion = fmt.Sprintf("http://localhost:%s", defaultPort)
		}

		prompt := fmt.Sprintf("Enter URL for %s", env)
		if suggestion != "" {
			prompt += fmt.Sprintf(" [%s]", suggestion)
		}
		fmt.Printf("%s: ", prompt)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" && suggestion != "" {
			input = suggestion
		}

		if input != "" {
			cfg.Envs[env] = input
		}
	}

	// 2. Save configuration
	if err := SaveConfig(LocalConfigPath, cfg); err != nil {
		fmt.Printf("❌ Error: Failed to save config: %v\n", err)
		return
	}

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
