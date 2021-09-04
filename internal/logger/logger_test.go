// Copyright 2021 Harness, Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package logger

import (
	"context"
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	l := NewLogger(true)

	ctx := WithContext(context.Background(), l)
	got := FromContext(ctx)

	if got != l {
		t.Errorf("Expected Logger from context")
	}
}

func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())
	if got != L {
		t.Errorf("Expected default Logger from context")
	}
}

func TestRequest(t *testing.T) {
	l := NewLogger(true)

	ctx := WithContext(context.Background(), l)
	req := new(http.Request)
	req = req.WithContext(ctx)

	got := FromRequest(req)

	if got != l {
		t.Errorf("Expected Logger from http.Request")
	}
}
