// Package sstable implements the on-disk sorted-string table format: block
// builder/reader, index and metaindex blocks, properties block, footer,
// table writer, and table reader with a two-level iterator.
//
// Depends on base and bloom.
package sstable
