# SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
# SPDX-License-Identifier: MIT

# Built by GoReleaser (dockers_v2): binaries come prebuilt in the context.
FROM gcr.io/distroless/static
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/testbed /usr/bin/testbed
ENTRYPOINT ["/usr/bin/testbed"]
