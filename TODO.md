# Updated Core Functionality

## install "edge" and "stable"

- install "edge" version from "develop" branch, by default
- "stable" version can be installed with a "--stable" option from the "main" branch

## Automatic Aliasing

The tool internally maps common shorthand to standard keys so you don't have to remember exactly how you saved it:
prod, p, live → production
stg, stage → staging
d, local → dev

## The "Missing Config" Assistant

If you run gourl prod in a project that hasn't been configured yet:
It checks for .cache/gourls.json.
If missing, it says: "No URLs configured for this project. Run 'gourl set prod ' to get started."
It checks if .cache is git-ignored and suggests: git config --global core.excludesfile ~/.gitignore_global && echo ".cache/" >> ~/.gitignore_global.

## Sample code

```go
package main

import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "runtime"
  "strings"
)

const configPath = ".cache/gourls.json"

type Config map[string]string

func main() {
  args := os.Args[1:]

  // 1. Default 'gourl' -> Open production
  if len(args) == 0 {
    openUrl("production")
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
  }
}

func listConfig() {
  cfg := loadConfig()
  if len(cfg) == 0 {
    fmt.Println("No URLs configured.")
    return
  }
  for env, url := range cfg {
    fmt.Printf("%-12s: %s\n", env, url)
  }
}

func printHelp() {
  fmt.Println("gourl - Project URL Manager")
  fmt.Println("\nUsage:")
  fmt.Println("  gourl              Opens 'production' URL")
  fmt.Println("  gourl <env>        Opens specific URL (e.g., prod, staging, dev)")
  fmt.Println("  gourl set <env> <url>   Save a URL")
  fmt.Println("  gourl list         List all saved URLs")
}

```

## purge

- [ ] "gourl --purge" should uninstall itself and remove all app files it has installed. ".cache/" folders will survive the purge

## interactive dev setup

- [ ] "interactive dev setup": gourl -i or --interactive should ask user interactively if they want to setup dev, prod, staging URLs, if the current project is an app being developed with source code present at the root of the current folder.

## env category

- [ ] "gourl set --global rails:dev <url>" should save the URL globally for rails dev category. "--global" is optional.
- [ ] "gourl dev" in a rails project folder should pickup "localhost:3000" from the standard global dev config list even if ".cache/urls.json" is not defined in that rails project folder.
- [ ] similarly for other common or popular languages like go, rust, node, python, etc.
- [ ] "gourl list" should list all saved URLs, grouped by category.

## global URLs

- [ ] "gourl set --global <env> <url>" should save the URL globally, not just for the current project.
- [ ] "gourl list --global" should list all globally saved URLs.
- [ ] "gourl dev" in a rails project folder should pickup "localhost:3000" from the standard global dev config list even if ".cache/urls.json" is not defined in that rails project folder.
- [ ] when installing, gourl should install global defaults for rails, go, rust, node, python and other common development languages. So, when user runs "gourl dev" in a rails project folder, it should pickup "localhost:3000" from the standard global dev config list even if ".cache/urls.json" is not defined in that rails project folder. Similarly for other languages.

## favourite URLs

- [ ] "gourl set --favourite <env> <url>" should save the URL as a favourite, not just for the current project.
