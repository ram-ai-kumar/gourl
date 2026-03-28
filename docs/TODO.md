# gourl TODO

## Release v0.2.0 (High Priority - P0)

### [x] unset

- **Command**: `gourl unset <env>`
- **Behavior**: Remove the URL from the environment mapping.
- **Scope**: Supports local and `--global` removal.

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

### [x] env category

- **Command**: `gourl set --global <env> <url>`
- **Behavior**: Save the URL globally for a specific project category (e.g., `rails`, `go`, `node`).
- **Standard Defaults**: Pre-seed global list with standard URLs during the first global set.
- **Listing**: `gourl list` groups and shows the source (local, global, or default).

### [x] global URLs

- **Command**: `gourl set --global <env> <url>`
- **Behavior**: Save the URL globally, accessible from any project folder.
- **Listing**: `gourl list` includes user mappings from `~/.cache/gourls.json`.

---

## Release v0.4.0 (Low Priority - P2)

### [ ] favourite URLs

- **Command**: `gourl set --favourite <env> <url>`
- **Behavior**: Save the URL as a favourite for quick access across different contexts.
