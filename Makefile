# ---------------------------------------------------------------------------- #
#              The MIT License (MIT) Copyright © 2023 The CUC Authors          #
#                                                                              #
#                +--------------------------------------------+                #
#                |                                            |                #
#                |           ██████╗██╗   ██╗ ██████╗         |                #
#                |          ██╔════╝██║   ██║██╔════╝         |                #
#                |          ██║     ██║   ██║██║              |                #
#                |          ██║     ██║   ██║██║              |                #
#                |          ╚██████╗╚██████╔╝╚██████╗         |                #
#                |           ╚═════╝ ╚═════╝  ╚═════╝         |                #
#                |                                            |                #
#                +--------------------------------------------+                #
#                                                                              #
#                CUC - Command-line URL Checker (and notifier)                 #
#                                                                              #
# ---------------------------------------------------------------------------- #
#                                                                              #
# Copyright © 2023 David Aparicio david.aparicio@free.fr                       #
#                                                                              #
# Permission is hereby granted, free of charge, to any person obtaining a copy #
# of this software and associated documentation files (the "Software"), to deal#
# in the Software without restriction, including without limitation the rights #
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell    #
# copies of the Software, and to permit persons to whom the Software is        #
# furnished to do so, subject to the following conditions:                     #
#                                                                              #
# The above copyright notice and this permission notice shall be included in   #
# all copies or substantial portions of the Software.                          #
#                                                                              #
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR   #
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,     #
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE  #
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER       #
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,#
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN    #
# THE SOFTWARE.                                                                #
#                                                                              #
# ---------------------------------------------------------------------------- #

GORELEASER_FLAGS ?= --snapshot --clean --rm-dist
all: compile check-format lint test

# Variables and Settings
version     ?=  $(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))# 0.0.1
target      ?=  cuc
org         ?=  davidaparicio
authorname  ?=  David Aparicio # The CUC Authors
authoremail ?=  david.aparicio@free.fr
license     ?=  MIT
year        ?=  2023
copyright   ?=  Copyright (c) $(year)

COMMIT      := $(shell git rev-parse HEAD)
DATE        := $(shell date)## +%Y-%m-%d)
PKG_LDFLAGS := github.com/davidaparicio/cuc/internal

CGO_ENABLED := 1
export CGO_ENABLED

compile: mod ## Compile for the local architecture ⚙
	@echo "Compiling..."
	go build -ldflags "\
	-s -w \
	-X '${PKG_LDFLAGS}.Version=$(version)' \
	-X '${PKG_LDFLAGS}.BuildDate=$(DATE)' \
	-X '${PKG_LDFLAGS}.Revision=$(COMMIT)'" \
	-o bin/$(target) .

# -X 'main.Version=$(version)' \
# -X 'main.AuthorName=$(authorname)' \
# -X 'main.AuthorEmail=$(authoremail)' \
# -X 'main.Copyright=$(copyright)' \
# -X 'main.License=$(license)' \
# -X 'main.Name=$(target)' \

.PHONY: run
run: ## Run the command with default values
	@echo "Running...\n"
	@go run -ldflags "\
	-s -w \
	-X '${PKG_LDFLAGS}.Version=$(version)' \
	-X '${PKG_LDFLAGS}.BuildDate=$(DATE)' \
	-X '${PKG_LDFLAGS}.Revision=$(COMMIT)'" \
	./main.go --url "http://neverssl.com/makeSSLgreatAgain" -c 404 -f "assets/mp3/ubuntu_desktop_login.mp3"

.PHONY: version
version: ## Run the command to get the version
	@echo "Running to get the command version...\n"
	@go run -ldflags "\
	-s -w \
	-X '${PKG_LDFLAGS}.Version=$(version)' \
	-X '${PKG_LDFLAGS}.BuildDate=$(DATE)' \
	-X '${PKG_LDFLAGS}.Revision=$(COMMIT)'" \
	./main.go version

.PHONY: goreleaser
goreleaser: ## Run goreleaser directly at the pinned version 🛠
	go run github.com/goreleaser/goreleaser@v1.18.2 $(GORELEASER_FLAGS)

.PHONY: mod
mod: ## Go mod things
# go mod tidy
# go get -d ./...

.PHONY: install
install: compile test ## Install the program to /usr/bin 🎉
	@echo "Installing..."
	sudo cp bin/$(target) /usr/local/bin/$(target)

.PHONY: test
test: compile ## 🤓 Run go tests
	@echo "Testing..."
	go test -v ./...

.PHONY: clean
clean: ## Clean your artifacts 🧼
	@echo "Cleaning..."
	rm -rvf dist/*
	rm -rvf release/*
	rm -rvf pkg/api/

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: format
format: ## Format the code using gofmt
	@echo "Formatting..."
	@gofmt -s -w $(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: check-format
check-format: ## Used by CI to check if code is formatted
	@gofmt -l $(shell find . -name '*.go' -not -path "./vendor/*") | grep ".*" ; if [ $$? -eq 0 ]; then exit 1; fi

.PHONY: lint
lint: ## Runs the linter
	golangci-lint run

.PHONY: check-editorconfig
check-editorconfig: ## Use to check if the codebase follows editorconfig rules
	@docker run --rm --volume=$(shell PWD):/check mstruebing/editorconfig-checker

.PHONY: doc
doc: ## Launch the offline Go documentation 📚
	@echo "open http://127.0.0.1:6060 and run godoc server..."
	open "http://127.0.0.1:6060"
	godoc -http=:6060 -play

.PHONY: fuzz
fuzz: ## Run fuzzing tests 🌀
	@echo "Fuzzing..."
#	go test -v -fuzz "Fuzz" -fuzztime 15s

.PHONY: benchmark
benchmark: ## Run benchmark tests 🚄
	@echo "Benchmarking..."
	go test -v -run=^$ -bench . -benchmem -benchtime=10s ./

.PHONY: sec
sec: ## Golang Security checks code for security problems 🔒
	gosec ./...
	govulncheck ./...