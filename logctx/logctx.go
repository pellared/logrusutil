// Package logctx adds contextual logging to logrus.
package logctx

import (
	"context"

	"github.com/sirupsen/logrus"
)

// DefaultLogEntry is used to create a new log entry if there is none in the context.
var DefaultLogEntry = logrus.NewEntry(logrus.StandardLogger())

type contextKey struct{}

// New creates a new context with log entry.
func New(ctx context.Context, logEntry *logrus.Entry) context.Context {
	return context.WithValue(ctx, contextKey{}, logEntry)
}

// From returns the log entry from the context.
// Returns log entry from DefaultLogEntry if there is no log entry in the context.
func From(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value(contextKey{}).(*logrus.Entry); ok {
		return entry
	}
	return DefaultLogEntry
}
