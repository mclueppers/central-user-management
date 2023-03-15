NAME = cum
PROJECT_DIR = $(shell pwd)
SRC_DIR = src/cum
VERSION = 0.1.0

.PHONY = build clean download

build:
	@echo "Building $(NAME) $(VERSION)..."
	cd $(SRC_DIR) && go build \
		-ldflags " \
			-X main.Version=$(VERSION) \
			-X main.BuildDate=$(shell date -u '+%Y-%m-%dT%I:%M:%SZ') \
			-X main.GitCommit=$(shell git rev-parse HEAD) \
			-X main.GitBranch=$(shell git rev-parse --abbrev-ref HEAD) \
			-X main.GitState=\"$(shell git status --porcelain)\" \
			-X main.GitSummary=$(shell git describe --tags --always --dirty) \
			-X main.GoVersion=\"$(shell go version)\" \
			" \
		-o $(PROJECT_DIR)/bin/$(NAME) \
		cum/cum/cmd

clean:
	@echo "Cleaning..."
	@rm -f $(NAME)

download:
	@echo "Downloading dependencies..."
	@cd $(SRC_DIR) && go get -d -v ./...
