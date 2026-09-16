// Package memtable implements the in-memory sorted write buffer: the
// skiplist and the memtable wrapper (size accounting, point Get, internal
// iterator) built on top of it.
//
// Depends only on base.
package memtable
