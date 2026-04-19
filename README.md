# gourl

A smart CLI utility for managing and quickly opening project URLs across different environments.

[![Go Report Card](https://goreportcard.com/badge/github.com/ram-ai-kumar/gourl)](https://goreportcard.com/report/github.com/ram-ai-kumar/gourl)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Security: Governance-Ready](https://img.shields.io/badge/Security-Governance--Ready-blue.svg)

## Key Features

- **🚀 Smart Command Routing**: Opens environment-specific URLs (`prod`, `stg`, `dev`) in your default browser.
- **🔗 Environment Aliases**: Use intuitive shorthand like `p`, `stg`, or `local`.
- **📂 Local Configuration**: Project-specific URL mappings stored in `.cache/gourls.json`.
- **💻 Cross-Platform**: Native support for macOS, Windows (CMD/PowerShell), and Linux (including WSL).
- **🛠️ Assisted Setup**: Helpful prompts and `.gitignore` suggestions for first-time users, with automatic detection for Go, Node.js (npm, Yarn, pnpm, Bun), Rails, Python, and Rust.
- **🛡️ Governance-Ready**: Secure-by-design, local-first architecture with zero external dependencies.

---

## Installation

### Option 1: Install Script (Recommended)

Default (**Edge** version from `develop` branch):

```bash
curl -sSL https://raw.githubusercontent.com/ram-ai-kumar/gourl/develop/scripts/install-source.sh | bash
```

Stable (**Stable** version from `main` branch):

```bash
curl -sSL https://raw.githubusercontent.com/ram-ai-kumar/gourl/main/scripts/install.sh | bash -s -- --stable
```

> [!NOTE]
> The modern installer includes **automatic source build fallback**. If a pre-compiled binary is not found for your specific architecture, it will automatically attempt to build from source using [Go](https://go.dev/doc/install).

### Other Options

- **Homebrew**: `brew install --head https://github.com/ram-ai-kumar/gourl/raw/develop/scripts/gourl.rb`
- **Go Install**: `go install github.com/ram-ai-kumar/gourl@latest`

---

## Creating a Release

To create a new release of gourl:

1. **Create a git tag** with the version number:

   ```bash
   git tag v0.2.1
   git push origin v0.2.1
   ```

2. **Automatic Release Process**:
   - GitHub Actions triggers the [`release.yml`](.github/workflows/release.yml) workflow
   - GoReleaser builds binaries for all platforms (Linux, macOS, Windows, FreeBSD, OpenBSD)
   - Version is injected via ldflags from the git tag
   - SHA256 checksums are generated for all artifacts
   - Release notes are automatically generated from git commits
   - Homebrew formula is pushed to `ram-ai-kumar/homebrew-tap`
   - Artifacts are uploaded to the GitHub Release

3. **Release Artifacts** (available in the release):
   - `gourl_v{version}_linux_amd64.tar.gz`
   - `gourl_v{version}_linux_arm64.tar.gz`
   - `gourl_v{version}_darwin_amd64.tar.gz`
   - `gourl_v{version}_darwin_arm64.tar.gz`
   - `gourl_v{version}_windows_amd64.zip`
   - `gourl_v{version}_freebsd_amd64.tar.gz`
   - `gourl_v{version}_freebsd_arm64.tar.gz`
   - `gourl_v{version}_openbsd_amd64.tar.gz`
   - `SHA256SUMS.txt`

---

## Local Testing

To test a release locally before pushing:

```bash
# Test with snapshot mode (dry-run)
goreleaser release --snapshot --clean

# Check artifacts in ./dist folder
ls -la dist/
```

To verify the version is correctly injected:

```bash
./dist/gourl_v{version}_linux_amd64 --version
```

---

## Documentation

For detailed guides and governance information, please explore our documentation suite:

- 📖 **[Usage Guide](docs/usage.md)**: Detailed command references, aliasing logic, and configuration formats.
- 🧪 **[Testing Guide](docs/testing.md)**: How we verify our enterprise security posture and feature correctness.
- ⚖️ **[Security & Governance](docs/governance.md)**: Executive summary of our security posture, including:
  - [Security Deep-dive](docs/security.md)
  - [Compliance & Supply Chain](docs/compliance.md)
  - [Zero Trust Architecture (ZTA)](docs/zta.md)

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
