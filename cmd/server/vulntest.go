// Copyright © 2026 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

// Calls a known-vulnerable symbol (GO-2020-0017) so govulncheck reports a
// called vulnerability. Testing only for OpenCHAMI/github-actions#53; revert.
package main

import jwt "github.com/dgrijalva/jwt-go"

func init() {
	_ = jwt.MapClaims{}.VerifyAudience("", false)
}
