# ADR-0001: Development Environment

## Context
The development machine runs Windows, but fsync semantics, directory fsync, `kill -9` behavior, and the Go race detector all behave most naturally on Linux. This project depends heavily on correct fsync/crash behavior, so the dev loop needs to exercise that behavior accurately.

## Options considered
- Windows native development.
- WSL2 (Ubuntu) with the repo on the Linux filesystem.
- Dual-boot Linux.

## Decision
Use WSL2 (Ubuntu LTS) as the primary development environment, with the repo living on the WSL Linux filesystem (`/home/...`), not under `/mnt/c`. Windows-native behavior is validated separately in CI (test matrix, see 0.4.2).

## Consequences
- Accurate Linux fsync/crash semantics during day-to-day development.
- `/mnt/c` would add 9P-layer I/O overhead and break fsync accuracy for both benchmarks and crash tests — avoided by keeping the repo under `/home/...`.
- Windows-specific differences the code must account for, exercised by the Windows CI job on every PR:
  - Directory fsync is a no-op on Windows.
  - An open file cannot be deleted on Windows (deletion is deferred) — this makes refcounting bugs surface immediately on the Windows job, which is treated as a feature, not something to skip.
  - The file-locking API differs from POSIX `flock`.

## Revisit when
If development moves to a team primarily on Windows-native tooling, or WSL2's fsync/crash behavior is found to diverge meaningfully from bare-metal Linux.
