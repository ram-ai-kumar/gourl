package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ram-ai-kumar/gourl/internal"
)

var Version = "dev"

func main() {
	// Root flags
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version information")
	purgeFlag := flag.Bool("purge", false, "Uninstall gourl")

	flag.Usage = internal.PrintHelp
	flag.Parse()

	if *helpFlag {
		internal.PrintHelp()
		return
	}
	if *versionFlag {
		fmt.Printf("gourl version %s\n", Version)
		return
	}
	if *purgeFlag {
		internal.PurgeApp()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		internal.PrintHelp()
		return
	}

	command := args[0]

	switch command {
	case "set":
		setCmd := flag.NewFlagSet("set", flag.ExitOnError)
		global := setCmd.Bool("global", false, "Save to global config")
		setCmd.BoolVar(global, "g", false, "Save to global config")
		favourite := setCmd.Bool("favourite", false, "Save to favourites")
		setCmd.BoolVar(favourite, "f", false, "Save to favourites")

		setCmd.Parse(args[1:])
		setArgs := setCmd.Args()

		if len(setArgs) < 2 {
			fmt.Println("Usage: gourl set [--global|--favourite] <env> <url>")
			os.Exit(1)
		}
		env := setArgs[0]
		url := setArgs[1]

		var path string
		var scope string

		if *favourite {
			var err error
			path, err = internal.GetFavouritesConfigPath()
			if err != nil {
				fmt.Printf("❌ Error: Could not determine favourites config path: %v\n", err)
				os.Exit(1)
			}
			scope = "favourite"
		} else if *global {
			var err error
			path, err = internal.GetGlobalConfigPath()
			if err != nil {
				fmt.Printf("❌ Error: Could not determine global config path: %v\n", err)
				os.Exit(1)
			}
			scope = "global"
		} else {
			path = internal.LocalConfigPath
			scope = "local"
		}

		cfg, _ := internal.LoadConfig(path)
		if *global {
			internal.PreseedDefaults(cfg)
		}
		if *favourite {
			cfg.Favourites[internal.NormalizeEnv(env)] = url
		} else {
			cfg.Envs[internal.NormalizeEnv(env)] = url
		}

		if err := internal.SaveConfig(path, cfg); err != nil {
			fmt.Printf("❌ Error: Failed to save config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Saved %s -> %s (%s)\n", internal.NormalizeEnv(env), url, scope)

	case "unset":
		unsetCmd := flag.NewFlagSet("unset", flag.ExitOnError)
		global := unsetCmd.Bool("global", false, "Remove from global config")
		unsetCmd.BoolVar(global, "g", false, "Remove from global config")
		favourite := unsetCmd.Bool("favourite", false, "Remove from favourites")
		unsetCmd.BoolVar(favourite, "f", false, "Remove from favourites")

		unsetCmd.Parse(args[1:])
		unsetArgs := unsetCmd.Args()

		if len(unsetArgs) < 1 {
			fmt.Println("Usage: gourl unset [--global|--favourite] <env>")
			os.Exit(1)
		}
		env := unsetArgs[0]
		
		var scope string
		var err error
		
		if *favourite {
			err = internal.UnsetFavourite(env)
			scope = "favourite"
		} else {
			err = internal.UnsetConfig(env, *global)
			scope = "local"
			if *global {
				scope = "global"
			}
		}
		
		if err != nil {
			fmt.Printf("❌ Error: Failed to unset environment: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ Unset %s (%s)\n", internal.NormalizeEnv(env), scope)

	case "list":
		internal.ListConfig()

	case "version", "-v":
		fmt.Printf("gourl version %s\n", Version)

	case "help", "-h":
		internal.PrintHelp()

	case "-i", "--interactive":
		internal.RunInteractiveSetup()

	default:
		// 2. 'gourl <env>' -> Open specific env
		openUrl(command)
	}
}

func openUrl(env string) {
	url, ok := internal.GetConfigValue(env)

	if !ok {
		fmt.Printf("❌ No URL found for '%s'. Run: gourl set %s <url>\n", env, env)

		// Check if local config file exists and suggest gitignore
		if _, err := os.Stat(internal.LocalConfigPath); os.IsNotExist(err) {
			fmt.Println("💡 No URLs configured for this project.")
			internal.CheckAndSuggestGitignore()
		}
		os.Exit(1)
	}

	internal.OpenBrowser(url)
}
