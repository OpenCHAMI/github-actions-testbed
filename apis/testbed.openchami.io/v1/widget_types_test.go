// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
// SPDX-License-Identifier: MIT

package v1

import (
	"context"
	"testing"

	"github.com/openchami/fabrica/pkg/fabrica"
)

func TestWidget(t *testing.T) {
	w := &Widget{Metadata: fabrica.Metadata{Name: "w1", UID: "uid-1"}}
	if got := w.GetKind(); got != "Widget" {
		t.Errorf("GetKind() = %q, want %q", got, "Widget")
	}
	if got := w.GetName(); got != "w1" {
		t.Errorf("GetName() = %q, want %q", got, "w1")
	}
	if got := w.GetUID(); got != "uid-1" {
		t.Errorf("GetUID() = %q, want %q", got, "uid-1")
	}
	if err := w.Validate(context.Background()); err != nil {
		t.Errorf("Validate() = %v, want nil", err)
	}
}
