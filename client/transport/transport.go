// Package transport re-exports internal/client/transport for external consumers.
package transport

import (
	itransport "github.com/distribution/distribution/v3/internal/client/transport"
)

// Types
type RequestModifier = itransport.RequestModifier

// Functions
var NewTransport = itransport.NewTransport
