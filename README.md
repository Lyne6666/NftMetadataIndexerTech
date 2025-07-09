# NftMetadataIndexerTech

## Description

A smart contract suite implementing a fractionalized NFT ownership model using ERC-1155 for shared governance and ERC-721 for the underlying asset, optimized for gas efficiency through bitwise operations on ownership rights.

## Features

- Indexes NFT metadata by parsing ERC-721 and ERC-1155 contract events directly from blockchain nodes.
- Stores NFT metadata in a normalized, schema-agnostic database optimized for complex queries and aggregations.
- Exposes a GraphQL API for efficient querying of NFT metadata, including support for filtering, sorting, and pagination.
- Implements a caching layer using Redis to minimize database load and improve API response times.
- Automatically detects and processes newly deployed NFT contracts using bloom filter subscription to event logs.
- Supports custom metadata parsers via a plugin architecture, allowing for indexing of non-standard NFT contracts.
- Provides a configurable retry mechanism for failed metadata indexing jobs, ensuring data consistency.
- Monitors data pipeline health using Prometheus metrics, including indexing latency and error rates.
## Installation

```bash
pip install git+https://github.com/Lyne6666/NftMetadataIndexerTech.git
```

## Usage

```bash
python -m nftmetadataindexertech --verbose
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
