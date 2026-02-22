# ADR 005: P2P PBFT Consensus Simulation

## Status
Proposed

## Context
We need to simulate multi-node block validation to test supply chain integrity in a distributed environment.

## Decision
Implement a simplified PBFT (Practical Byzantine Fault Tolerance) algorithm in the Go backend. This includes PRE-PREPARE, PREPARE, and COMMIT phases with a basic quorum logic.

## Consequences
- Allows testing of network latency and node failures.
- Increases complexity of the local development environment.
