// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
// SPDX-License-Identifier: MIT

// Command testbed is a trivial binary used to exercise the
// OpenCHAMI/github-actions build and release workflows.
package main

import "fmt"

// Set via -ldflags by GoReleaser.
var (
	version   = "dev"
	commit    = "none"
	buildHost = "unknown"
)

// Info returns the build metadata line printed by the binary.
func Info() string {
	return fmt.Sprintf("testbed %s (commit %s, built on %s)", version, commit, buildHost)
}

func main() {
	fmt.Println(Info())
}
