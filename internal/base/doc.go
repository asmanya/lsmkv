// Package base holds the foundational types shared across the storage
// engine: InternalKey, Kind, SeqNum, FileNum, the filename scheme, error
// types, and the InternalIterator interface.
//
// It sits at the bottom of the dependency graph and imports nothing but the
// standard library, so that memtable, sstable, and iterator can each depend
// on base without depending on one another.
package base
