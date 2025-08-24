# Rust Launcher

This is the Rust implementation of the **PlayHarbor** utility.  
It provides the same functionality as the Go version:

- Launch an Epic Games Store title using its launch URL.
- Wait for the game process to start.
- Keep running until the game process exits.

The Rust version is designed with future extensibility in mind.

## Usage

The launcher expects **two arguments**:
1. The Epic Games Store launch URL (e.g., `com.epicgames.launcher://apps/...`)
2. The executable name of the game (without `.exe` extension)

### Example
```bash
playdock.exe "com.epicgames.launcher://apps/Fortnite?action=launch" "FortniteClient-Win64-Shipping"
```

The program will:
- Open the Epic Games URL (Epic Launcher handles the game start)
- Wait 5 seconds
- Search for the game process
- Block until the game process exits

## Build

Make sure you have [Rust installed](https://www.rust-lang.org/).

From this directory:

```bash
cargo build --release
```

The compiled binary will be in:

```
target/release/playdock.exe
```

## Dependencies

This implementation uses:
- [sysinfo](https://crates.io/crates/sysinfo) for process management.

Dependencies are declared in `Cargo.toml` and managed automatically by Cargo.
