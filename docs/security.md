# Security Posture Statement

`gourl` is engineered to provide local URL management with an emphasis on data integrity and environment isolation.

## Data Isolation & Local-First Philosophy

`gourl` follows a **Zero-Cloud, Zero-Telemetry** policy. Your data is your property and stays where you put it.

- **Local Storage**: All URL mappings are stored in `.cache/gourls.json` within your project directories.
- **No Remote Calls**: `gourl` never sends your URL data to any remote server or analytics platform. The only outgoing network calls are made by the **Installer** (to fetch the latest releases from GitHub).
- **Process Isolation**: Every project folder maintain its own `.cache` folder, preventing cross-project data leakage.

### Local Data Flow
```text
[ Developer CLI ] --(1) Read config)--> [ .cache/gourls.json ]
        |                                       |
        |                                (Local Project Scope)
        |                                       |
        +-------(2) System Open)-------> [ Default Browser ]
```

## Minimized Attack Surface

### Zero External Dependencies
`gourl` is built using only the **Go Standard Library**. This significantly reduces the supply chain risk and eliminates the possibility of "left-pad" style library hijackings or hidden vulnerabilities in third-party packages. This includes our interactive framework detection logic (`gourl -i`), which resolves project types using native filesystem APIs.

### Principle of Least Privilege
`gourl` does not require root or `sudo` permissions for normal operation. It runs entirely in the user's workspace, reducing the potential impact of any local exploit. The `purge` command ensures that the binary can be removed by the user without elevated permissions, maintaining a minimal system footprint.

## Binary Integrity & Distribution

### Secure Download
The **Installer (`install.sh`)** uses TLS (HTTPS) exclusively for all downloads from GitHub's official release channels. 

### Modern Build Pipeline
Our build system uses [Go Releaser](https://goreleaser.com/) to ensure clean, reproducible binaries from the verified source code on branch `main`.

## Threat Modeling

| Threat Vector | Mitigation Strategy |
| ------------- | ------------------- |
| **URL Injection** | `gourl` uses the system's native `exec.Command` and standard browser openers (`open`, `xdg-open`, `rundll32`), providing robust platform-level sanitization. |
| **Config Tampering** | Config files are stored in the user's workspace with standard filesystem permissions. Users are encouraged to gitignore `.cache/`. |
| **Long-Term Persistence** | The `purge` feature allows users to reliably uninstall the tool, mitigating risks associated with outdated or forgotten binaries on a system. |
| **Supply Chain** | Zero-external-dependency posture minimizes this risk to the core Go Standard Library itself. |
