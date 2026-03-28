# Security & Governance Portfolio

`gourl` is engineered with a "Security-First, Privacy-Always" philosophy. This document provides an executive summary of our governance standards and security posture.

## Executive Summary

As a developer tool designed for local project management, `gourl` adheres to strict data minimization and local-first architectural patterns. We believe that your project URLs are sensitive metadata and should never leave your local machine without your explicit action.

## Core Governance Pillars

### 🔐 [Security Posture](file:///Users/ram/Work/code/dev-stack/gourl/docs/security.md)
Detailed analysis of how `gourl` handles data and system interactions.
- **Local-Only Storage**: All configuration is stored locally in `.cache/gourls.json`.
- **Zero-Dependency Architecture**: Minimized attack surface using only the Go Standard Library.
- **Secure Distribution**: Binary integrity and verified release channels.

### 🛡️ [Compliance & Supply Chain](file:///Users/ram/Work/code/dev-stack/gourl/docs/compliance.md)
Information on software integrity and regulatory alignment.
- **SBOM (Software Bill of Materials)**: Full visibility into the build components.
- **Vulnerability Management (SLA)**: Proactive monitoring and patching strategy for the Go toolchain.
- **Licensing**: Open-source transparency under the MIT license.

### 🌐 [Zero Trust Architecture (ZTA)](file:///Users/ram/Work/code/dev-stack/gourl/docs/zta.md)
Alignment with modern Zero Trust principles.
- **Never Trust, Always Verify**: How URLs are validated before browser invocation.
- **Principle of Least Privilege**: Running entirely in user-space without elevated permissions.
- **Environment Isolation**: Logical separation of project-level and global configurations.

---

## Contact & Maintenance

For security disclosures or governance inquiries, please refer to the official repository at [ram-ai-kumar/gourl](https://github.com/ram-ai-kumar/gourl).
