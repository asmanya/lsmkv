// Package tablecache implements a refcounted LRU cache of open sstable
// readers, so that hot tables don't get reopened on every access.
//
// Depends on base, sstable, and vfs.
package tablecache
