# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOBUILDPLUGIN=$(GOBUILD) -buildmode=plugin
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOGENERATE=$(GOCMD) generate
GOLIST=$(GOCMD) list

# Docker parameters
DOCKERCMD=docker
DOCKEREXEC=$(DOCKERCMD) exec -it

DOCKERCCMD=docker-compose
DOCKERCUP=$(DOCKERCCMD) up
DOCKERCDOWN=$(DOCKERCCMD) down

# Sources parameters
SOURCE_ENTRYPOINT=./cmd/main.go

# Binary parameters
BINARY_NAME=app
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Targets
help:
	@echo "Compilation, build, run"
	@echo "build               : Build the app"
	@echo "test                : Launch all units tests"
	@echo "test_cover          : Launch all units tests with the cover of each package"
	@echo "clean               : Remove temporary files"

build:
		env GOOS=linux GOARCH=386 $(GOBUILD) -o $(BINARY_PATH).out -v $(SOURCE_ENTRYPOINT)
test:
		$(GOTEST) -v ./...
test_cover:
		$(GOTEST) -v ./... -coverprofile=coverage.txt -covermode=atomic
run:
	@echo "Run the program"
	$(GORUN) $(SOURCE_ENTRYPOINT)
clean:
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
dockerc_up:
		$(DOCKERCUP) --build
dockerc_down:
		$(DOCKERCDOWN)
docker_exec_db:
		$(DOCKEREXEC) db bash

