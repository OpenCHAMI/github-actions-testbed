// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
// SPDX-License-Identifier: MIT

package main

import (
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	if got := Info(); !strings.HasPrefix(got, "testbed ") {
		t.Fatalf("Info() = %q, want prefix %q", got, "testbed ")
	}
}
