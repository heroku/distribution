// Package auth re-exports internal/client/auth for external consumers.
package auth

import (
	"context"

	iauth "github.com/distribution/distribution/v3/internal/client/auth"
	"github.com/distribution/distribution/v3/internal/client/auth/challenge"
	"github.com/distribution/distribution/v3/internal/client/transport"
)

// Types
type AuthenticationHandler = iauth.AuthenticationHandler
type CredentialStore = iauth.CredentialStore

// Errors
var ErrNoBasicAuthCredentials = iauth.ErrNoBasicAuthCredentials

// Functions
var NewAuthorizer = iauth.NewAuthorizer
var NewBasicHandler = iauth.NewBasicHandler

// targetAppKey is the context key for TargetApp.
type targetAppKey struct{}

// WithTargetApp returns a new context with the target app identity stored.
func WithTargetApp(ctx context.Context, app string) context.Context {
	return context.WithValue(ctx, targetAppKey{}, app)
}

// TargetApp extracts the target app identity from the context.
// Returns the app string and true if present, or empty string and false otherwise.
func TargetApp(ctx context.Context) (string, bool) {
	app, ok := ctx.Value(targetAppKey{}).(string)
	return app, ok
}

// Re-export types from sub-packages that are commonly used with auth.
// These are here for convenience but consumers should prefer importing
// the sub-packages directly.

// ChallengeManager is an alias for challenge.Manager.
type ChallengeManager = challenge.Manager

// RequestModifier is an alias for transport.RequestModifier.
type RequestModifier = transport.RequestModifier
