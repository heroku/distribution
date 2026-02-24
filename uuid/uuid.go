// Package uuid re-exports internal/uuid for external consumers.
package uuid

import (
	iuuid "github.com/distribution/distribution/v3/internal/uuid"
)

// Loggerf is a function that can be set to enable warning logging for UUID
// generation issues. By default, it is a no-op.
var Loggerf func(string, ...interface{}) = func(string, ...interface{}) {}

// NewString returns a new UUID string.
var NewString = iuuid.NewString
