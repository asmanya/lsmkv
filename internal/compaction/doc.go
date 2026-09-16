// Package compaction implements the size-tiered compaction picker (a pure
// function over a Version) and the compaction iterator's drop rules
// (shadowed versions, tombstones).
//
// Depends on base, manifest, iterator, and sstable.
package compaction
