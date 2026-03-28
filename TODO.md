# gourl TODO

## Release v0.2.0 (High Priority - P0)

### [ ] purge
- **Command**: `gourl --purge`
- **Behavior**: Should uninstall itself and remove all app files it has installed.
- **Persistence**: `.cache/` folders will survive the purge to ensure user configurations are not lost if they reinstall later.
- **Scope**: Remove binary from `~/.local/bin/` and any associated global config files (once implemented).

### [ ] add godog (cucumber framework)
- **Goal**: Full test coverage using the godog framework.
- **Status**: Framework added, but needs comprehensive test cases defined to cover entire functionality including edge cases and platform-specific behaviors.

### [ ] build and release
- **Task**: Write a script/feature to build and release the next version of this tool.
- **Integration**: Should handle cross-platform builds and potentially GitHub Release uploads.

---

## Release v0.2.1 (Medium Priority - P1)

### [x] interactive dev setup
- **Command**: `gourl -i` or `--interactive`
- **Behavior**: Ask user interactively if they want to setup `dev`, `prod`, `staging` URLs.
- **Context Awareness**: Should detect if the current project is an app being developed (e.g., check for source code presence like `go.mod`, `package.json`, etc. at the root).

---

## Release v0.3.0 (Medium Priority - P1)

### [ ] env category
- **Command**: `gourl set --global rails:dev <url>`
- **Behavior**: Save the URL globally for a specific project category (e.g., `rails`, `go`, `node`).
- **Flags**: `--global` is optional if a category prefix is used.
- **Smart Resolution**: `gourl dev` in a rails project folder should pickup `localhost:3000` from the standard global `rails:dev` config even if `.cache/gourls.json` is missing.
- **Coverage**: Support common languages like Go, Rust, Node, Python, etc.
- **Listing**: `gourl list` should group results by category.

### [ ] global URLs
- **Command**: `gourl set --global <env> <url>`
- **Behavior**: Save the URL globally, accessible from any project folder.
- **Listing**: `gourl list --global` to show global mappings.
- **Default Integration**: During installation, `gourl` should pre-seed global defaults for common frameworks (e.g., `rails` -> `localhost:3000`, `go` -> `localhost:8080`, etc.).

---

## Release v0.4.0 (Low Priority - P2)

### [ ] favourite URLs
- **Command**: `gourl set --favourite <env> <url>`
- **Behavior**: Save the URL as a favourite for quick access across different contexts.
