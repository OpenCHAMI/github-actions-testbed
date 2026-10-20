# SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
# SPDX-License-Identifier: MIT

IMAGE_TAG   ?= $(shell git describe --tags --always --abbrev=0)
QUADLET_DIR ?= build/quadlet
CLEANFILES  := bin/ build/ dist/
# cmd/client is Fabrica's generated CLI; only the server ships.
BINARIES    := testbed:./cmd/server

include Makefile.common

.PHONY: quadlet-render
quadlet-render: ## Pin the quadlet image tag for packaging (accepts IMAGE_TAG, QUADLET_DIR)
	mkdir -p $(QUADLET_DIR)
	if ! grep -q '@IMAGE_TAG@' packaging/systemd/testbed.container; then \
		echo "error: @IMAGE_TAG@ placeholder missing from testbed.container" >&2; \
		exit 1; \
	fi
	sed 's|@IMAGE_TAG@|$(IMAGE_TAG)|' packaging/systemd/testbed.container > $(QUADLET_DIR)/testbed.container
