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

### Saving a URL
```bash
gourl set <env> <url>
```
Example: `gourl set staging https://staging.myapp.com`

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
Project-specific configuration is stored in:
`.cache/gourls.json`

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

- **Cross-Platform**: Native support for macOS (`open`), Windows (`rundll32`), and Linux (`xdg-open` or `wslview` for WSL).
- **Assisted Setup**: Helpful prompts when environments are missing or configuration is incomplete.
- **Robust Installation**: The modern installer supports binary downloads with automatic source build fallback if the binary is unavailable for your architecture.
