# gourl

A smart CLI utility for managing and quickly opening project URLs across different environments.

## Features

- **Smart Command Routing** - Opens production URL by default, or specific environments
- **Environment Aliases** - Use shorthand like `prod`, `p`, `stg`, `d` instead of full names
- **Local Configuration** - Project-specific URL mappings stored in `.cache/gourls.json`
- **Cross-platform Support** - Works on macOS, Windows, and Linux (including WSL)
- **Missing Config Assistant** - Helpful setup guidance when first used

## Installation

### Option 1: Install Script (Recommended)

```bash
curl -sSL https://raw.githubusercontent.com/ram-ai-kumar/gourl/main/install-source.sh | bash
```

### Option 2: Homebrew

```bash
brew tap ram-ai-kumar/gourl
brew install ram-ai-kumar/gourl/gourl
```

### Option 3: Go Install

```bash
go install github.com/ram-ai-kumar/gourl@latest
```

### Option 4: Build from Source

```bash
git clone https://github.com/ram-ai-kumar/gourl.git
cd gourl
go build -o gourl
```

Or run directly:

```bash
go run main.go
```

## Usage

### Basic Commands

```bash
# Show help message
gourl

# Open specific environment
gourl prod          # Opens production URL
gourl staging       # Opens staging URL
gourl dev           # Opens development URL

# Save a URL for an environment
gourl set prod https://myapp.com
gourl set dev http://localhost:3000
gourl set api https://api.myapp.com

# List all configured URLs
gourl list

# Show help
gourl help
```

### Environment Aliases

The tool automatically maps common shorthand to standard environment names:

| Aliases             | Maps to      |
| ------------------- | ------------ |
| `prod`, `p`, `live` | `production` |
| `stg`, `stage`      | `staging`    |
| `d`, `local`        | `dev`        |

### First-time Setup

When you first run `gourl` in a new project, it will guide you through setup:

```bash
$ gourl
❌ No URL found for 'production'. Run: gourl set production <url>
💡 No URLs configured for this project.
💡 To ignore .cache/ in git, run:
   echo '.cache/' >> .gitignore
```

## Configuration

URLs are stored in `.cache/gourls.json` in your project directory:

```json
{
  "production": "https://myapp.com",
  "staging": "https://staging.myapp.com",
  "dev": "http://localhost:3000",
  "api": "https://api.myapp.com"
}
```

## Example Workflow

```bash
# Initial setup
gourl set prod https://myapp.com
gourl set dev http://localhost:3000

# Daily usage
gourl prod          # Opens production in browser
gourl dev           # Opens local dev server
gourl               # Shows help message

# Check what's configured
gourl list
```

## Building

```bash
go build -o gourl
```

This creates a `gourl` executable you can move to your PATH for system-wide use.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
