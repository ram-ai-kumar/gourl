# Zero Trust Architecture (ZTA) Posture Statement

`gourl` adheres to modern "Zero Trust" principles for its application-level security and local environment isolation.

## The Three Pillars of ZTA for gourl

### 1. **Never Trust, Always Verify**
`gourl` does not assume its environment's security.

- **URL Validation**: All incoming requests to open a URL must match a pre-verified configuration in the project's local JSON file.
- **Context-Aware Verification**: The tool's guided setup (`gourl -i`) detects the underlying project architecture (e.g., Go, Node, Rails) to provide verified environment defaults, reducing configuration errors.
- **System Separation**: `gourl` separates the responsibility of *storing* a URL from the responsibility of *opening* it by leveraging the user's default, system-hardened browser environment.

### 2. **Principle of Least Privilege (PoLP)**
- **User-Space Operation**: `gourl` is designed to be installed and executed entirely in user-space (`~/.local/bin/`).
- **No Elevated Privileges**: The tool never requests `sudo` or root permissions, limiting its potential blast radius.

### 3. **Environment Isolation**
- **Logical Air-Gapping**: Project-specific configurations are physically and logically segregated into individual `.cache` folders.
- **Cross-Contamination Protection**: The tool lacks a central "shared" configuration by default (unless specifically implemented as a global feature), preventing a breach in one project folder from automatically compromising URLs in another.
- **Ephemeral System Footprint**: The `purge` command allows for complete removal of the tool's binary and system-level installation artifacts, ensuring a clean state when it's no longer required.

### Environment Isolation Model
```text
/Project-A/             /Project-B/             /Project-C/
  |                       |                       |
  +-- .cache/gourls.json  +-- .cache/gourls.json  +-- .cache/gourls.json
  | (Isolated Data)       | (Isolated Data)       | (Isolated Data)
  |                       |                       |
  V                       V                       V
[gourl CLI]             [gourl CLI]             [gourl CLI]
```

---

## Technical ZTA Implementation

- **Identity & Access Management (IAM)**: `gourl` relies on the host operating system's filesystem permissions for all its ACLs.
- **Network Security**: By not having any remote dependencies or telemetry, `gourl` maintains a zero-network-trust state during runtime.

---

## Vision for Future ZTA Enhancements

As `gourl` evolves, we plan to implement additional ZTA-aligned features:
- **Project-Specific Sandboxing**: Enhanced checks to ensure an environment's URLs are only accessible when the user is within the specific project directory.
- **Configuration Checksums**: Cryptographic verification of `.cache/gourls.json` to detect unauthorized tampering.
- **Secure Keyring Integration**: Evaluating options for storing sensitive URLs in the system's native keychain (e.g., macOS Keychain or Windows Credential Manager).
