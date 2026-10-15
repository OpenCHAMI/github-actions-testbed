#!/bin/sh
# SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
# SPDX-License-Identifier: MIT
#
# Shared by the rpm (%preun) and deb (prerm) packages. Stop the service only
# on real removal: rpm passes 0, deb passes "remove".

case "$1" in
    0|remove)
        systemctl stop testbed.service >/dev/null 2>&1 || :
        ;;
esac
