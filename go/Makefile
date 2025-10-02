SHELL := /bin/bash

BINARY_NAME := delivrio
BUILD_FLAGS :=

GIT_TAG := $(shell git describe --tags --abbrev=0)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%:z)

# Default target to build the binary
.PHONY: build
build:
	go build -o ${BINARY_NAME} -ldflags "-X main.AppVersion=${GIT_TAG} -X main.BuildTime=${BUILD_TIME}" ${BUILD_FLAGS}

CONFIG_FILE=delivrio_config.yaml

checkconfig:
	@if [ ! -f $(CONFIG_FILE) ]; then \
		echo "Config file $(CONFIG_FILE) not found."; \
		read -p "Do you want to create it? (y/n) " yn; \
		case $$yn in \
			[Yy]*) $(MAKE) delivrio_config.yaml;; \
			[Nn]*) echo "Please create the config file and rerun the make command."; exit 1;; \
		esac; \
	else \
		echo "Config file $(CONFIG_FILE) found."; \
	fi

# Local development target
.PHONY: local
local: build
	./${BINARY_NAME}

# Local new target with demo data, run "local" subsequently
.PHONY: localnew
localnew: build
	${MAKE} checkconfig
	./${BINARY_NAME} -seed=true

delivrio_config.yaml:
	./${BINARY_NAME} -create-config=true

# Desktop build target with desktop tags
.PHONY: desktop
desktop:
	${MAKE} build BUILD_FLAGS="-tags desktop"
	./${BINARY_NAME}

.PHONY: build-review
build-review:
	${MAKE} build BUILD_FLAGS="-tags internal"

.PHONY: build-production
build-production:
	${MAKE} build BUILD_FLAGS="-tags internal"

# Count lines of code
.PHONY: loc
loc:
	find . -type f -name '*.go' \
      ! -path './ent/*' \
      ! -path './gengql/generated/*' \
      -exec cat {} + | wc -l

# Run tests with coverage, excluding specific directories
.PHONY: test
test:
	go test -failfast -cover $(shell go list ./... | grep -v /ent/)

# Default to canary, specify version manually: curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.24.0 sh
# Run full migration for SQLite and Postgres
.PHONY: migrate
migrate: migrate-sqlite migrate-postgres

.PHONY: migrate-sqlite
migrate-sqlite:
	atlas migrate diff $(name).up \
		--dir "file://ent/migrate/migrations/sqlite" \
		--to "ent://schema" \
		--dev-url "sqlite://file?mode=memory&_fk=1" || exit 1

.PHONY: migrate-postgres
migrate-postgres:
	# Sudo since docker seems to require it; might be a misconfigured install
	sudo atlas migrate diff $(name).up \
		--dir "file://ent/migrate/migrations/postgres" \
		--to "ent://schema" \
		--dev-url "docker://postgres/16/delivrio?search_path=public" || exit 1
	sudo chmod 777 -R ent/migrate/migrations/postgres || exit 1

# Hash migrations
.PHONY: migrate-hash
migrate-hash:
	atlas migrate hash --dir "file://ent/migrate/migrations/postgres"
	atlas migrate hash --dir "file://ent/migrate/migrations/sqlite"

# Lint migrations
.PHONY: migrate-lint
migrate-lint:
	atlas migrate lint \
		--dir "file://ent/migrate/migrations/sqlite" \
		--dev-url "sqlite://file?mode=memory&_fk=1" \
		--latest 10
	atlas migrate lint \
		--dir "file://ent/migrate/migrations/postgres" \
		--dev-url "docker://postgres/16/delivrio?search_path=public" \
		--latest 10

# Local development target for internal development
.PHONY: localinternal
localinternal: build
	./${BINARY_NAME} -config=../../delivrio-internal/config/delivrio_config.yaml

# Local new target with seed flag and internal users
.PHONY: localnew
localnewinternal:
	${MAKE} build BUILD_FLAGS="-tags internal"
	./${BINARY_NAME} -seed=true -config=../../delivrio-internal/config/delivrio_config.yaml
