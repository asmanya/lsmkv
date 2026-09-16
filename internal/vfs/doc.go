// Package vfs defines the filesystem interface all storage-engine I/O goes
// through: the OS implementation, a strict in-memory MemFS for power-loss
// simulation, a FaultFS wrapper for fault injection, and directory locking.
//
// No package outside vfs's own OS implementation may call os.* directly;
// everything else uses this interface so that filesystem behavior can be
// simulated and faults injected in tests. Imports stdlib plus
// golang.org/x/sys.
package vfs
