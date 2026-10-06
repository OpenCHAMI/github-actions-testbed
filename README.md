<!--
SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
SPDX-License-Identifier: MIT
-->

# github-actions-testbed

[![Latest release](https://img.shields.io/github/v/release/OpenCHAMI/github-actions-testbed)](https://github.com/OpenCHAMI/github-actions-testbed/releases/latest)
[![Test](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/test.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/test.yml)
[![Coverage](https://coveralls.io/repos/github/OpenCHAMI/github-actions-testbed/badge.svg?branch=main)](https://coveralls.io/github/OpenCHAMI/github-actions-testbed?branch=main)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/OpenCHAMI/github-actions-testbed/badge)](https://scorecard.dev/viewer/?uri=github.com/OpenCHAMI/github-actions-testbed)

<details>
<summary>Additional project checks</summary>

**Build quality**

[![Release with GoReleaser](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/release.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/release.yml)
[![Build](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/build.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/build.yml)
[![Lint](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/lint.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/lint.yml)
[![REUSE compliance check](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/reuse.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/reuse.yml)

**Security**

[![CodeQL](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/github-code-scanning/codeql)
[![Vulnerability Check](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/govulncheck.yml/badge.svg)](https://github.com/OpenCHAMI/github-actions-testbed/actions/workflows/govulncheck.yml)

</details>

A project that tests the
[OpenCHAMI/github-actions](https://github.com/OpenCHAMI/github-actions) reusable
workflows end to end, releases included, and doubles as the template for wiring
them into a service repo. Nothing depends on its images, packages, or releases.

It exercises every GoReleaser release effect: binaries, archives, a multi-arch
image (`dockers_v2`), rpm/deb quadlet packages, signing (with a throwaway key
generated per run), attestations, and the GitHub release.

## Using this as a template

1. Copy `.github/workflows/`, `.goreleaser.yaml`, `Dockerfile`, `Makefile`'s
   `quadlet-render` target, and `packaging/`.
2. Rename `testbed` / `github-actions-testbed` to your project (binary, image,
   package name, quadlet, URLs, badges).
3. Update the expected package file lists in `build.yml` and `release.yml`
   (`validate-packages`): rpm lists every file and owned directory, deb only
   files.
4. Pin `OpenCHAMI/github-actions` to a release (or SHA) instead of the testbed's
   branch pin.
5. Replace the throwaway `test-key` job in `release.yml` with your real signing
   key secret. Annotated GoReleaser reference:
   [`.goreleaser.example.yml`](https://github.com/OpenCHAMI/github-actions/blob/main/.goreleaser.example.yml).

## Testing the workflows

- **PR mode**: open a PR. `build.yml` builds and validates packages;
  `release.yml` pushes `ghcr.io/openchami/github-actions-testbed:pr-<N>`;
  closing the PR runs `cleanup.yml`.
- **Release mode**: push a `v*` tag (e.g. `v0.0.1`). Tag freely.
- **Another github-actions ref**: update the pinned SHA in
  `.github/workflows/*.yml`.
