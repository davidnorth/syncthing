// Copyright (C) 2025 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !windows
// +build !windows

package main

import (
	"github.com/syncthing/syncthing/lib/config"
	"github.com/syncthing/syncthing/lib/syncthing"
)

// initSystray is a stub for non-Windows platforms
func initSystray(app *syncthing.App, cfg config.Wrapper) {
	// No-op on non-Windows platforms
}

// shouldRunSystray returns false on non-Windows platforms
func shouldRunSystray() bool {
	return false
}
