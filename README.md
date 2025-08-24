
[![CI](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml/badge.svg)](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.22-blue.svg)](https://golang.org/doc/go1.22)
[![Rust Version](https://img.shields.io/badge/rust-stable-orange.svg)](https://www.rust-lang.org/)
[![Go Coverage](https://img.shields.io/badge/go--coverage-pending-lightgrey.svg)](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml)
[![Rust Coverage](https://img.shields.io/badge/rust--coverage-pending-lightgrey.svg)](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml)
[![Go Lint](https://img.shields.io/badge/go--lint-passing-brightgreen.svg)](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml)
[![Rust Lint](https://img.shields.io/badge/rust--lint-passing-brightgreen.svg)](https://github.com/feed3r/play-harbor/actions/workflows/ci.yml)

# PlayHarbor

This project provides simple utilities to launch Epic Games Store titles
in a way that is also visible to Steam and Steam Link.

Currently, two implementations are available:
- `go-launcher/`: lightweight Go version.
- `rust-launcher/`: Rust version with future extensibility.

The long-term goal is to create a more complete application
to manage Epic Games Store titles and improve integration with Steam and other game
management softwares.

## Roadmap
- [x] Initial launcher in Go
- [x] Initial launcher in Rust
- [ ] Steam overlay detection
- [ ] Epic games library parsing
- [ ] Unified launcher with GUI

## Contributing

We welcome contributions from everyone!

Before contributing, please read our [Contributing Guide](CONTRIBUTING.md) which explains:
- How to submit issues and pull requests
- Branching strategy
- Commit message conventions (we follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/))
- Code style and testing requirements

## License
This project is licensed under the [Apache License 2.0](LICENSE).
