# Security & Governance Portfolio

`gourl` is engineered with a "Security-First, Privacy-Always" philosophy. This document provides an executive summary of our governance standards and security posture.

## Executive Summary

`gourl` (Project URL Manager) is a secure, local-first utility for developers to manage environment-specific URLs efficiently. Built on a foundation of **Zero-Cloud, Zero-Telemetry**, the tool ensures that sensitive internal URLs never leave the developer's local machine.

### Key Governance Values
- **Data Sovereignty**: Project configurations remain in local `.cache/` folders.
- **Context Awareness**: Guided setup (`gourl -i`) uses filesystem markers to suggest safe environment defaults.
- **Clean Disposal**: Binary uninstallation (`gourl --purge`) leaves no persistent executable footprint behind.
- **Auditability**: Zero external dependencies and MIT-licensed source code for complete transparency.

### 🔐 [Security Posture](security.md)
Detailed analysis of how `gourl` handles data and system interactions.
- **Local-Only Storage**: All configuration is stored locally in `.cache/gourls.json`.
- **Zero-Dependency Architecture**: Minimized attack surface using only the Go Standard Library.
- **Secure Distribution**: Binary integrity and verified release channels.

### 🛡️ [Compliance & Supply Chain](compliance.md)
Information on software integrity and regulatory alignment.
- **SBOM (Software Bill of Materials)**: Full visibility into the build components.
- **Vulnerability Management (SLA)**: Proactive monitoring and patching strategy for the Go toolchain.
- **Licensing**: Open-source transparency under the MIT license.

### 🌐 [Zero Trust Architecture (ZTA)](zta.md)
Alignment with modern Zero Trust principles.
- **Never Trust, Always Verify**: How URLs are validated before browser invocation.
- **Principle of Least Privilege**: Running entirely in user-space without elevated permissions.
- **Environment Isolation**: Logical separation of project-level and global configurations.

---

## Contact & Maintenance

For security disclosures or governance inquiries, please refer to the official repository at [ram-ai-kumar/gourl](https://github.com/ram-ai-kumar/gourl).
