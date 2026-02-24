// Package dcontext re-exports internal/dcontext for external consumers.
package dcontext

import (
	"context"
	"net/http"

	idcontext "github.com/distribution/distribution/v3/internal/dcontext"
)

// Types
type Logger = idcontext.Logger

// Functions
var (
	WithRequest         = idcontext.WithRequest
	GetRequestID        = idcontext.GetRequestID
	WithResponseWriter  = idcontext.WithResponseWriter
	GetResponseWriter   = idcontext.GetResponseWriter
	GetRequestLogger    = idcontext.GetRequestLogger
	GetResponseLogger   = idcontext.GetResponseLogger
	WithVars            = idcontext.WithVars
	WithLogger          = idcontext.WithLogger
	GetLoggerWithField  = idcontext.GetLoggerWithField
	GetLoggerWithFields = idcontext.GetLoggerWithFields
	GetLogger           = idcontext.GetLogger
)

// Errors
var (
	ErrNoRequestContext        = idcontext.ErrNoRequestContext
	ErrNoResponseWriterContext = idcontext.ErrNoResponseWriterContext
)

// GetRequest extracts the *http.Request from the context. Returns
// ErrNoRequestContext if the request is not available.
func GetRequest(ctx context.Context) (*http.Request, error) {
	v := ctx.Value("http.request")
	r, ok := v.(*http.Request)
	if !ok || r == nil {
		return nil, ErrNoRequestContext
	}
	return r, nil
}
