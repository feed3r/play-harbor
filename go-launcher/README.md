# Go Launcher

This is the Go implementation of the **PlayHarbor** utility.  
It allows you to start an Epic Games Store title using its launch URL, 
wait for the game process to start, and keep running until the game exits.  
This is useful when launching Epic games through Steam or Steam Link.

## Usage

The launcher expects **two arguments**:
1. The Epic Games Store launch URL (e.g. `com.epicgames.launcher://apps/...`)
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

Make sure you have [Go installed](https://go.dev/dl/).

From this directory:

```bash
go build -o playdock.exe main.go
```

This will produce a standalone `playdock.exe` binary.

## Dependencies

This implementation uses:
- [gopsutil](https://github.com/shirou/gopsutil) for process management.

Dependencies are declared in `go.mod` and will be downloaded automatically by Go.
