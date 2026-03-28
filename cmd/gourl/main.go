package main

import (
	"fmt"
	"os"

	"github.com/ram-ai-kumar/gourl/internal"
)

var Version = "dev"

func main() {
	args := os.Args[1:]

	// 1. Default 'gourl' -> Show help
	if len(args) == 0 {
		internal.PrintHelp()
		return
	}

	command := args[0]

	switch command {
	case "set":
		if len(args) < 3 {
			fmt.Println("Usage: gourl set <env> <url>")
			os.Exit(1)
		}
		if err := internal.SaveConfig(args[1], args[2]); err != nil {
			fmt.Printf("❌ Error: Failed to save config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Saved %s -> %s\n", internal.NormalizeEnv(args[1]), args[2])
	case "list":
		internal.ListConfig()
	case "version", "--version", "-v":
		fmt.Printf("gourl version %s\n", Version)
	case "help", "--help", "-h":
		internal.PrintHelp()
	case "-i", "--interactive":
		internal.RunInteractiveSetup()
	case "--purge":
		internal.PurgeApp()
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
			internal.PurgeApp()
			return
		}
		// 2. 'gourl <env>' -> Open specific env
		openUrl(command)
	}
}

func openUrl(env string) {
	env = internal.NormalizeEnv(env)
	cfg := internal.LoadConfig()
	url, ok := cfg[env]

	if !ok {
		fmt.Printf("❌ No URL found for '%s'. Run: gourl set %s <url>\n", env, env)

		// Check if config file exists and suggest gitignore
		if _, err := os.Stat(internal.ConfigFileName); os.IsNotExist(err) {
			fmt.Println("💡 No URLs configured for this project.")
			internal.CheckAndSuggestGitignore()
		}
		os.Exit(1)
	}

	internal.OpenBrowser(url)
}
