// Package record implements the log record framing shared by the WAL and
// the MANIFEST: a writer and reader that operate on io.Writer/io.Reader,
// with checksums and error classification for torn tails vs corruption.
//
// It only imports the standard library, so both manifest and the WAL code
// can depend on it without pulling in vfs directly.
package record
