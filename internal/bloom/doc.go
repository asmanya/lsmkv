// Package bloom implements a deterministic Bloom filter: parameter math,
// hash function, builder, and reader.
//
// Imports only the standard library, so sstable can depend on it without
// pulling in anything else.
package bloom
