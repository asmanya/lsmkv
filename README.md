# lsmkv

A from-scratch LSM-tree storage engine in Go, built for durability under crashes and correctness under concurrency.

**Status:** under construction.

[![CI](https://github.com/asmanya/lsmkv/actions/workflows/ci.yml/badge.svg)](https://github.com/asmanya/lsmkv/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/asmanya/lsmkv.svg)](https://pkg.go.dev/github.com/asmanya/lsmkv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Planned features

- Write-ahead log with crash-safe recovery
- In-memory memtable (skiplist) backed by the WAL
- Sorted-string table (SSTable) on-disk format with block-level checksums
- Size-tiered compaction with tombstone garbage collection
- Bloom filters for point-lookup acceleration
- Fault-injection and crash-recovery test harness
- Benchmark suite with statistically validated results (benchstat)

## License

MIT — see [LICENSE](LICENSE).
