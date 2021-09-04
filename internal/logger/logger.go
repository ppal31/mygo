// Copyright 2021 Harness, Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.
package logger

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/log"
)

// L is an alias for the logger with debug set to false as default
var L = NewLogger(false)

type loggerKey struct{}

// Info extracts the logger from context and writes a structured info leve log
func Info(ctx context.Context, msg string, fields ...interface{}) {
	level.Info(FromContext(ctx)).Log(append([]interface{}{"msg", msg}, fields...)...)
}

// Debug extracts the logger from context and writes a structured debug level log
func Debug(ctx context.Context, msg string, fields ...interface{}) {
	level.Debug(FromContext(ctx)).Log(append([]interface{}{"msg", msg}, fields...)...)
}

// Error extracts the logger from context and writes a structured error level log
func Error(ctx context.Context, msg string, fields ...interface{}) {
	level.Error(FromContext(ctx)).Log(append([]interface{}{"msg", msg}, fields...)...)
}

func WithFields(ctx context.Context, fields ...interface{}) context.Context {
	return WithContext(ctx, log.WithPrefix(FromContext(ctx), fields...))
}

// WithContext returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithContext(ctx context.Context, l log.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

// FromContext retrieves the current logger from the context. If no
// logge is available, the default logger is returned.
func FromContext(ctx context.Context) log.Logger {
	l := ctx.Value(loggerKey{})
	if l == nil {
		return L
	}
	return l.(log.Logger)
}

// FromRequest retrieves the current logger from the request. If no
// logger is available, the default logger is returned.
func FromRequest(r *http.Request) log.Logger {
	return FromContext(r.Context())
}

func NewLogger(debug bool) log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	allowedLevels := level.AllowInfo()
	if debug {
		allowedLevels = level.AllowDebug()
	}
	logger = level.NewFilter(logger, allowedLevels)
	logger = log.With(
		logger,
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(5),
	)

	return logger
}
