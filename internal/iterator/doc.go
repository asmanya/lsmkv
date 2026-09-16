// Package iterator implements the heap-based merging iterator that combines
// multiple InternalIterator sources (memtables, sstables) into one sorted
// stream, with tie-breaking and error propagation.
//
// Depends only on base, whose InternalIterator interface it implements and
// consumes.
package iterator
