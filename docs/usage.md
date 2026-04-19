# gourl Usage Guide

`gourl` is a smart CLI utility for managing and quickly opening project URLs across different environments.

## Basic Commands

### Show Help
```bash
gourl
# or
gourl help
```

### Opening a URL
```bash
gourl <env>
```
Example: `gourl prod` opens the production URL in your default browser.

Example: `gourl set staging https://staging.myapp.com`

#### Save a Global URL
```bash
gourl set --global <env> <url>
# or
gourl set -g <env> <url>
```
Global URLs are stored in `~/.cache/gourls.json` and serve as defaults across all projects.

### Removing a URL
```bash
gourl unset <env> [options]
```
Options:
- `--global`, `-g`: Remove from global configuration.
Example: `gourl unset staging` (removes from local config).

### Listing All URLs
```bash
gourl list
```
Displays all configured URLs for the current project.

### Check Version
```bash
gourl version
# or
gourl -v
```

### Uninstallation
```bash
gourl --purge
```
This command uninstalls `gourl` by instantly removing the binary from your system. 

*Note: Your project-specific `.cache/` folders will not be deleted.*

### Interactive Setup
```bash
gourl -i
# or
gourl --interactive
```
Launches a guided setup to configure `dev`, `staging`, and `production` URLs for the current project. 
`gourl` automatically detects common project types (Go, Node.js, Yarn, pnpm, Bun, Rails, Python, Rust) and suggests default development ports (e.g., `localhost:3000` for Node/Rails).

---

## Environment Aliases

`gourl` automatically maps common shorthand to standard environment names, so you don't have to worry about consistency:

| Aliases             | Internal Key |
| ------------------- | ------------ |
| `prod`, `p`, `live` | `production` |
| `stg`, `stage`      | `staging`    |
| `d`, `local`        | `dev`        |

You can use either the shorthand or the full name interchangeably.

---

## Configuration Details

### Storage Location
- **Local Config**: `.cache/gourls.json` (Project-specific)
- **Global Config**: `~/.cache/gourls.json` (User-specific defaults)

### Resolution Priority
When you run `gourl <env>`, the tool resolves the URL in the following order:
1.  **Local Project Config** (Overrides all)
2.  **Global User Config**
3.  **Project-Specific Defaults** (e.g., `8080` for Go, `3000` for Node, pre-seeded into the global config)

### Format
The configuration is a simple JSON mapping:
```json
{
  "production": "https://myapp.com",
  "staging": "https://staging.myapp.com",
  "dev": "http://localhost:3000"
}
```

### Git Integration
It is recommended to exclude the `.cache/` directory from version control. `gourl` will automatically remind you to do this if it detects a missing configuration:
```bash
echo '.cache/' >> .gitignore
```

---

## Technical Features

- **Missing Config Assistant**: If you run `gourl prod` in a project that hasn't been configured yet, the tool will:
    1. Check for the missing `.cache/gourls.json`.
    2. Prompt you with: `No URLs configured for this project. Run 'gourl set prod <url>' to get started.`
    3. Detect if `.cache/` is git-ignored and suggest the appropriate command to update your `.gitignore`.
- **Robust Installation**: The modern installer supports binary downloads with automatic source build fallback if the binary is unavailable for your architecture.

---

## Development & Testing

For contributors or developers looking to verify their installation:

### Unit Tests
```bash
make test
```

### Feature Tests (godog)
`gourl` uses the `godog` (Cucumber) framework for behavioral testing.
```bash
make test-features
```
*For a detailed breakdown of our test suite, see our [Testing Guide](testing.md).*
