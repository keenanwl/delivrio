//go:build linux

package main

import (
	"log"
	"os"
	"path/filepath"
)

func startOnLogin() error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	autostartDir := filepath.Join(homeDir, ".config", "autostart")
	err = os.MkdirAll(autostartDir, os.ModePerm)
	if err != nil {
		return err
	}

	autostartFile := filepath.Join(autostartDir, "DELIVRIO.desktop")

	if _, err := os.Stat(autostartFile); err == nil {
		log.Println("Autostart file already exists:", autostartFile)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	desktopEntry := `[Desktop Entry]
Type=Application
Exec=/usr/bin/delivrio
Hidden=false
NoDisplay=false
X-GNOME-Autostart-enabled=true
Name=DELIVRIO
Comment=Start DELIVRIO on login
`

	err = os.WriteFile(autostartFile, []byte(desktopEntry), 0644)
	if err != nil {
		return err
	}

	log.Println("Autostart file created:", autostartFile)
	return nil
}
