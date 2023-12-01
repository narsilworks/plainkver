// Package plainkver is an interface package to allow multiple implementations
//
// Author: Elizalde Baguinon
// Created: December 1, 2023
// Description: plainkver is an interface package to allow multiple implementations
// of plain key-value packages that uses different database client libraries
package plainkver

// PlainKVer is an interface for plain key-value packages.
// A plain key-value database makes use of existing database packages. It has no database of its own.
// The default table name for plain key-value should be called 'KeyValueTbl'.
// It can be changed by the implementing package.
type PlainKVer interface {
	// Open a connection to the database
	Open() error

	// Begin a transaction
	Begin() error

	// Get a value by key
	Get(key string) ([]byte, error)

	// Set the value of a key
	Set(key string, value []byte) error

	// Set the current bucket. The default bucket name is 'default'.
	SetBucket(bucket string) error

	// Delete a key
	Del(key string) error

	// Get the MIME of a key. It has its own bucket.
	GetMime(key string) (string, error)

	// Set the MIME of a key. It has its own bucket.
	SetMime(key string, mime string) error

	// List keys with the pattern.
	ListKeys(pattern string) ([]string, error)

	// Tally gets the current tally of a key.
	// A tally can be useful for a specific key if it
	// is used as a header of many details
	Tally(key string, offset int) (int, error)

	// Increment the key tally.
	TallyIncr(key string) (int, error)

	// Decrement the key tally.
	TallyDecr(key string) (int, error)

	// Reset the key tally.
	TallyReset(key string) error

	// Commit transaction
	Commit() error

	// Rollback transaction
	Rollback() error

	// Close the connection to the database
	Close() error
}
