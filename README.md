# Blockchain Supply Chain Integrity

High-integrity ledger system built with Go for industrial supply chains.

## Architecture
- **Merkle Tree**: Core validation structure.
- **Cryptographic Signing**: Ensures block authenticity.
- **Storage Adapters**: Pluggable storage for ledger data.

## SRE/Monitoring
- Ledger health checks and block propagation monitoring.
- Transaction failure rate tracking via metrics.

## ADR
- [ADR-001: SHA-256 for Hashing](docs/adr/001-sha256-hashing.md)
- [ADR-002: Modular Storage Adapter Pattern](docs/adr/002-modular-storage.md)
