# Compliance & Supply Chain Statement

`gourl` is committed to delivering software with high integrity, transparency, and a robust vulnerability management strategy.

## Software Bill of Materials (SBOM)

As a Go project with ZERO external dependencies (Standard Library only), our SBOM is minimalist by design.

**Component Status**:
- **Core Engine**: Go Standard Library (Standard library components are maintained and audited by the Go Security Team).
- **External Packages**: None.

## Vulnerability Management (SLA)

### Monitoring
We proactively monitor the following channels for vulnerabilities:
- **Go Security Announcements**: For any CVEs discovered in the Go standard library.
- **GitHub Security Advisories**: Automated Dependabot checks (even on zero-dependency projects) for underlying build toolchain issues.

### Patching SLA (Vulnerability Response)
In the event of a discovered vulnerability in the Go toolchain:
- **Triage**: Within 24 hours of public disclosure.
- **Remediation**: Build against the latest patched Go version.
- **Redistribution**: Immediate update to the official GitHub `main` branch, enabling all users to patch via `install.sh`.

### Verification Tools
- **CodeScan**: Regular use of `govulncheck` to identify known vulnerabilities in the current build environment.

---

## Data Sovereignty & Retention

`gourl` is engineered for maximum data control. All configuration data is stored locally within the user's project directories.

- **Storage Location**: `.cache/gourls.json`
- **Retention Policy**: Configuration data is retained indefinitely by default until manually deleted.
- **Uninstallation Behavior (Purge)**: Running `gourl --purge` removes the application's executable but **explicitly preserves** all project-specific `.cache/` directories. This ensures project continuity while allowing for tool-level cleanup.
- **Contextual Privacy**: The guided setup (`gourl -i`) uses only local filesystem markers (e.g., `go.mod`) to suggest non-identifiable, standard development ports.

---

## Licensing & IP

- **Transparency**: Fully open-source under the [MIT License](LICENSE).
- **Attribution**: No third-party components requiring attribution are currently part of the `gourl` core.

---

## Technical Auditing

Users are encouraged to audit the source code directly at [github.com/ram-ai-kumar/gourl](https://github.com/ram-ai-kumar/gourl). The `install.sh` script is fully scriptable and transparent for automated security reviews.
