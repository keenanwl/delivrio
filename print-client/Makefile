# Makefile

# Project name
PROJECT_NAME := delivrio-print

# Output directory
BUILD_DIR := build

# Go parameters
GO := go
GOOS := windows
GOARCH := amd64

# Cross-compilation variables
CC := x86_64-w64-mingw32-gcc
CXX := x86_64-w64-mingw32-g++
CGO_ENABLED := 1

# Target executable
TARGET := $(BUILD_DIR)/$(PROJECT_NAME).exe
CC := x86_64-w64-mingw32-gcc
CXX := x86_64-w64-mingw32-g++
GOOS := windows
GOARCH := amd64

# Source files
SRC := $(wildcard *.go)

# Version for Debian package
VERSION := 0.0.12

.PHONY: all clean build check-deps install-go install-mingw dev migrate migrate-hash build-debian

# Default target
all: check-deps build

# Ensure dependencies are installed
check-deps: install-go install-mingw

# Install Go if not already installed
install-go:
	@which go >/dev/null 2>&1 || { \
		echo "Installing Go..."; \
		wget -q https://dl.google.com/go/go1.16.7.linux-amd64.tar.gz; \
		sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz; \
		rm go1.16.7.linux-amd64.tar.gz; \
		export PATH=$$PATH:/usr/local/go/bin; \
	}

# Install mingw-w64 if not already installed
install-mingw:
	@which x86_64-w64-mingw32-gcc >/dev/null 2>&1 || { \
		echo "Installing mingw-w64..."; \
		sudo apt-get update; \
		sudo apt-get install -y mingw-w64; \
	}

# Clean build directory
clean:
	rm -rf $(BUILD_DIR)

# Development command with workaround for a known issue
dev:
	WEBKIT_DISABLE_COMPOSITING_MODE=1 wails dev -tags webkit2_41

# Migration commands
migrate:
	atlas migrate diff $(name).up \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "sqlite://file?mode=memory&_fk=1"

migrate-hash:
	atlas migrate hash --dir "file://ent/migrate/migrations"

# This doesn't work unless you install windows make
# only works on widnows. c&p to PS to run.
build-windows:
	wails build -nsis

# Build Debian package
build-debian:
	wails build
	fpm -s dir -t deb \
		-p "delivrio-$(VERSION)-1-amd64.deb" \
		-v "$(VERSION)" \
		--maintainer "Keenan <kwl@delivrio.com>" \
		--vendor "DELIVRIO" \
		--url "https://delivrio.com" \
		--description "The official DELIVRIO print client" \
		--architecture "amd64" \
		--name delivrio \
		--chdir $(BUILD_DIR) \
		bin/delivrio=/usr/bin/delivrio \
		linux/delivrio.1=/usr/share/man/man1/delivrio.1 \
		linux/delivrio.desktop=/usr/share/applications/delivrio.desktop \
		icons/hicolor/16x16/delivrio.png=/usr/share/icons/hicolor/16x16/apps/delivrio.png \
		icons/hicolor/24x24/delivrio.png=/usr/share/icons/hicolor/24x24/apps/delivrio.png \
		icons/hicolor/48x48/delivrio.png=/usr/share/icons/hicolor/48x48/apps/delivrio.png \
		icons/hicolor/64x64/delivrio.png=/usr/share/icons/hicolor/64x64/apps/delivrio.png \
		icons/hicolor/128x128/delivrio.png=/usr/share/icons/hicolor/128x128/apps/delivrio.png \
		icons/hicolor/256x256/delivrio.png=/usr/share/icons/hicolor/256x256/apps/delivrio.png \
		icons/hicolor/512x512/delivrio.png=/usr/share/icons/hicolor/512x512/apps/delivrio.png \
		icons/hicolor/1024x1024/delivrio.png=/usr/share/icons/hicolor/1024x1024/apps/delivrio.png

