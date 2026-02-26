// Package challenge re-exports internal/client/auth/challenge for external consumers.
package challenge

import (
	ichallenge "github.com/distribution/distribution/v3/internal/client/auth/challenge"
)

// Types
type Manager = ichallenge.Manager

// Functions
var NewSimpleManager = ichallenge.NewSimpleManager
