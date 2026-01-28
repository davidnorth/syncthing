// Copyright (C) 2025 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build windows
// +build windows

package main

import (
	_ "embed"
	"log/slog"

	"github.com/getlantern/systray"
	"github.com/syncthing/syncthing/internal/slogutil"
	"github.com/syncthing/syncthing/lib/config"
	"github.com/syncthing/syncthing/lib/svcutil"
	"github.com/syncthing/syncthing/lib/syncthing"
)

//go:embed assets/logo-32.png
var iconData []byte

// initSystray initializes the system tray icon and menu for Windows
func initSystray(app *syncthing.App, cfg config.Wrapper) {
	systray.Run(func() {
		onSystrayReady(app, cfg)
	}, onSystrayExit)
}

// onSystrayReady is called when the systray is ready
func onSystrayReady(app *syncthing.App, cfg config.Wrapper) {
	systray.SetIcon(iconData)
	systray.SetTitle("Syncthing")
	systray.SetTooltip("Syncthing - File Synchronization")

	// Create menu items
	mOpenUI := systray.AddMenuItem("Open Web UI", "Open Syncthing Web Interface")
	systray.AddSeparator()
	
	mPause := systray.AddMenuItem("Pause All", "Pause all devices and folders")
	mResume := systray.AddMenuItem("Resume All", "Resume all devices and folders")
	systray.AddSeparator()
	
	mRestart := systray.AddMenuItem("Restart", "Restart Syncthing")
	systray.AddSeparator()
	
	mExit := systray.AddMenuItem("Exit", "Exit Syncthing")

	// Handle menu clicks
	go func() {
		for {
			select {
			case <-mOpenUI.ClickedCh:
				if guiCfg := cfg.GUI(); guiCfg.Enabled {
					if err := openURL(guiCfg.URL()); err != nil {
						slog.Error("Failed to open Web UI", slogutil.Error(err))
					}
				}
			case <-mPause.ClickedCh:
				setPauseState(cfg, true)
				slog.Info("Paused all devices and folders via system tray")
			case <-mResume.ClickedCh:
				setPauseState(cfg, false)
				slog.Info("Resumed all devices and folders via system tray")
			case <-mRestart.ClickedCh:
				slog.Info("Restarting via system tray")
				app.Stop(svcutil.ExitRestart)
				return
			case <-mExit.ClickedCh:
				slog.Info("Exiting via system tray")
				app.Stop(svcutil.ExitSuccess)
				return
			}
		}
	}()
}

// onSystrayExit is called when the systray exits
func onSystrayExit() {
	// Cleanup if needed
}

// shouldRunSystray returns true if we should run the system tray
func shouldRunSystray() bool {
	// Run systray on Windows unless explicitly disabled
	// We can add a flag later if needed
	return true
}
