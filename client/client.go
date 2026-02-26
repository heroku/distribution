// Package client re-exports internal/client for external consumers.
package client

import (
	iclient "github.com/distribution/distribution/v3/internal/client"
)

// NewRepository creates a new Repository for the given repository name, base
// URL, and transport.
var NewRepository = iclient.NewRepository
