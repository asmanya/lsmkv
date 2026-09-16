package base

import "testing"

// TestSanity is a placeholder so CI has something to run against the
// skeleton before real base types (InternalKey, Kind, SeqNum, ...) land.
func TestSanity(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("arithmetic is broken")
	}
}
