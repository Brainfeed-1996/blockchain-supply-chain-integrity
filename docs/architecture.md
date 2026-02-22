# Blockchain Supply Chain Integrity

## Architecture

```mermaid
graph TD
    subgraph "Domain Layer"
        Block[Block Entity]
        Merkle[Merkle Tree]
    end

    subgraph "Service Layer"
        Ledger[Ledger Service]
    end

    subgraph "Adapter Layer"
        Storage[Storage Adapter]
        Crypto[Crypto Adapter]
        Memory[Memory Storage Impl]
    end

    Ledger --> Block
    Ledger --> Merkle
    Ledger --> Storage
    Storage --> Memory
```

### Layers

1.  **Domain**: Core business logic (Block, Merkle Tree).
2.  **Services**: Orchestrates operations (Ledger validation, block creation).
3.  **Adapters**: External interfaces (Storage, Cryptography).
