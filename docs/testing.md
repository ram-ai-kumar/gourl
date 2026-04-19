# gourl Testing Guide

`gourl` employs a comprehensive testing strategy to Silicon Valley and Enterprise standards, ensuring that every feature—especially security and compliance—is verified before release.

## Test Suite Overview

We use two primary testing tiers:

1. **Unit Tests**: Rapid validation of core logic (Go Standard Library).
2. **Feature Tests (BDD)**: Behavioral specifications using `godog` (Cucumber) to verify end-to-end scenarios, security posture, and zero-trust isolation.

---

## Running Tests

### Automated via Makefile

The simplest way to run all tests is via the provided `Makefile`:

```bash
# Run unit tests
make test

# Run feature tests (godog)
make test-features

# Run everything
make test-all
```

### Manual godog execution

If you have `godog` installed globally:

```bash
go test -v ./test/
```

---

## Feature Test Categories

Our `features/` directory contains specialized tests for every governance pillar:

| Category                          | Description                                                             |
| --------------------------------- | ----------------------------------------------------------------------- |
| **`smoke_tests.feature`**         | High-level "sanity checks" to ensure basic binary operations.           |
| **`basic_functionality.feature`** | Verifies `set`, `list`, and environment opening logic.                  |
| **`environment_aliases.feature`** | Ensures `prod`/`p` and `stg`/`stage` resolve correctly.                 |
| **`purge.feature`**               | Verifies self-uninstallation while preserving `.cache/` data.           |
| **`interactive.feature`**         | Tests the guided configuration and framework detection.                 |
| **`security_tests.feature`**      | Validates URL sanitization and process-level security.                  |
| **`security_tests.feature`**      | Validates URL sanitization and process-level security.                  |
| **`zta_tests.feature`**           | Verifies project-level isolation and "Never Trust" logic.               |
| **`compliance_tests.feature`**    | Checks for SBOM integrity and licensing.                                |
| **`negative_tests.feature`**      | Verifies graceful error handling for missing configs or invalid inputs. |
| **`edge_cases.feature`**          | Tests unusual inputs, long URLs, and platform-specific quirks.          |

---

## Development Setup

To begin contributing or running feature tests for the first time:

1. **Install godog**:

   ```bash
   make dev-setup
   # or
   go install github.com/cucumber/godog/cmd/godog@latest
   ```

2. **Run with Verbose Output**:
   ```bash
   go test -v ./test/ --godog.format=pretty
   ```

---

## CI/CD Integration

Every pull request to `main` or `develop` triggers the full test suite in GitHub Actions, ensuring zero regressions in our security posture.

---

## Specialized Test Logics

### Interactive CLI Simulation

To test the `gourl -i` guided setup, we use the `runGourlCommandWithInputs` helper in `test/godog_test.go`. This function:

1.  **Redirects Stdin**: Uses a Go pipe to feed simulated user responses into the tool.
2.  **Concurrency**: Spawns a goroutine to write inputs to the tool's stdin while the main process executes.
3.  **Captures Context**: Verifies that the tool correctly identifies the project type (e.g., `go.mod`) even during simulated interactions.

### Binary Self-Removal (Purge)

Testing `gourl --purge` requires the tool to delete its own binary. Our test suite handles this by:

1.  **Copying Binary**: Creating a temporary build of `gourl` in a sandbox directory.
2.  **Executing Sandbox**: Running the command on the sandboxed binary.
3.  **Post-Run Verification**: Confirming the sandbox file exists _before_ and is gone _after_ the command completes.
