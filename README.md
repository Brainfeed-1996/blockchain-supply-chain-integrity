# 🔗 Blockchain Supply Chain Integrity

[![Go](https://img.shields.io/badge/Language-Go-00ADD8.svg)](https://golang.org)
[![Blockchain](https://img.shields.io/badge/Tech-Blockchain-blue.svg)]()
[![Integrity](https://img.shields.io/badge/Status-Validated-brightgreen.svg)]()

---
Part of the [Industrial Portfolio 2026](https://github.com/Brainfeed-1996/industrial-portfolio-2026) ecosystem.

## Architecture
- **Merkle Tree**: Core validation structure.
- **ZKP Verification**: Zero-Knowledge Proof logic for privacy-preserving audits.
- **Cryptographic Signing**: Ensures block authenticity.
- **Storage Adapters**: Pluggable storage for ledger data.

## Deployment
Production-ready Docker configuration included.
- **Live Ledger:** [https://supply-chain.brainfeed.tech](https://supply-chain.brainfeed.tech)
- **Vercel Frontend:** [https://blockchain-sc-ui.vercel.app](https://blockchain-sc-ui.vercel.app)

## SRE/Monitoring
- Ledger health checks and block propagation monitoring.
- Transaction failure rate tracking via metrics.

## ADR
- [ADR-001: SHA-256 for Hashing](docs/adr/001-sha256-hashing.md)
- [ADR-002: Modular Storage Adapter Pattern](docs/adr/002-modular-storage.md)
