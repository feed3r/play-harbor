---
description: Enforce English as the only language for everything written in play-harbor's development surface — source code, comments, log/print/error strings, test names and assertion messages, commit messages, PR titles/bodies. Use before writing any of these, and when reviewing existing code, to catch and fix Italian text found along the way.
argument-hint: ""
allowed-tools: Read, Edit, Grep
---

# English-only development

**Rule:** everything written as part of developing play-harbor is in English,
with no exceptions — source code, comments, log/print/error strings, test
names and assertion messages, commit messages, PR titles and bodies.
Conversation with the user can stay in whatever language they use; this rule
only covers what gets written into the repository.

**Why:** explicit user instruction (2026-09-15) — Italian had crept into
freshly-written test assertion messages, and the user wants it caught and
fixed both going forward and retroactively wherever it's found.

## When writing new code

Before writing a comment, log line, error message, test name, or commit
message, check it's in English. Don't translate literally from an Italian
mental draft — write directly in idiomatic English.

## When reviewing/touching existing code

If a file you're already editing for another reason contains Italian text
(comments, string literals, test messages), fix it in the same change if the
fix is small and localized. Don't go out of scope hunting for it across the
whole repo unless asked — fix it "as you go," not as a dedicated sweep.

## Known pre-existing Italian text (fix opportunistically when touching these files)

- `go-launcher/runlauncher/runlauncher_test.go` — comments ("Mock delle
  dipendenze...", "Override le funzioni globali...", "Helper per creare...")
  and several `assert.*` message strings (e.g. "RunLauncher con manager
  attivo dovrebbe restituire nil", "errore lancio").
- `go-launcher/processutil/process_test.go` — several `assert.*` message
  strings (e.g. "errore inatteso", "pid inatteso", "atteso nessun errore",
  "expected nessun errore da Name").
- `go-launcher/config/config_test.go` — `require.NoError(t, err, "LoadConfig
  fallita")`.

Update this list when a listed item gets fixed, and add newly found items
here instead of fixing them silently with no record.
