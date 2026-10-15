# SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
# SPDX-License-Identifier: MIT

IMAGE_TAG   ?= $(shell git describe --tags --always --abbrev=0)
QUADLET_DIR ?= build/quadlet

.PHONY: all
all: build

.PHONY: build
build: ## Build the binary
	go build -o bin/testbed ./cmd/testbed

.PHONY: test
test: ## Run unit tests
	go test -race ./...

.PHONY: quadlet-render
quadlet-render: ## Pin the quadlet image tag for packaging (accepts IMAGE_TAG, QUADLET_DIR)
	mkdir -p $(QUADLET_DIR)
	if ! grep -q '@IMAGE_TAG@' packaging/systemd/testbed.container; then \
		echo "error: @IMAGE_TAG@ placeholder missing from testbed.container" >&2; \
		exit 1; \
	fi
	sed 's|@IMAGE_TAG@|$(IMAGE_TAG)|' packaging/systemd/testbed.container > $(QUADLET_DIR)/testbed.container

.PHONY: clean
clean: ## Remove build outputs
	rm -rf bin build dist
