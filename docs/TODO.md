# gourl TODO

## Release v0.5.0 (Medium Priority - P1)

### [ ] Cache folder organization

- **Issue**: Currently, configuration files are stored directly in `~/.cache/` as:
  - `~/.cache/gourls.json`
  - `~/.cache/gourl-favourites.json`
- **Requirement**: All gourl-related files and folders must be inside `~/.cache/gourl/`
- **New structure**:
  - `~/.cache/gourl/config.json` - Main configuration (local and global URLs)
  - `~/.cache/gourl/favourites.json` - Favourite URLs
  - `~/.cache/gourl/logs/` - Log files (future feature)
  - `~/.cache/gourl/cache/` - Temporary cache data (future feature)
- **Migration**: Implement automatic migration from old paths to new structure
- **Backward compatibility**: Support reading from old paths during transition period
- **Files to update**:
  - [`internal/config.go`](internal/config.go) - Update `GetGlobalConfigPath()` and `GetFavouritesConfigPath()`
  - [`internal/config.go`](internal/config.go:10) - Update `LocalConfigPath` constant
  - Documentation and README references

---
