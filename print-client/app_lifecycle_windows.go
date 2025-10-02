//go:build windows

package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"path/filepath"
	"runtime"
)

func startOnLogin() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("opening windows registry key: %w", err)
	}
	defer k.Close()

	basePath := `C:\Program Files (x86)`
	if is64BitOS() {
		basePath = `C:\Program Files`
	}

	appName := "DELIVRIO"
	execPath := filepath.Join(basePath, `DELIVRIO Print Client\DELIVRIO Print Client\DELIVRIO Print Client.exe`)

	err = k.SetStringValue(appName, execPath)
	if err != nil {
		return fmt.Errorf("setting %s to %q: %w", appName, execPath, err)
	}
	log.Printf("Added %s to start on login with path %q\n", appName, execPath)

	return nil
}

func is64BitOS() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}
