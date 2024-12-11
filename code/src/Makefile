GO ?= go
PROTOC ?= protoc

GOGO_PROTOBUF_PATH := $(shell go list -e -f '{{.Dir}}' github.com/gogo/protobuf/gogoproto)/../
PROTO_FILES=$(shell find . -type f -not -path "./vendor/*" -name "*.proto")
PWD=$(shell pwd)
DATE=$(shell date +%y%m%d)
REMOTE_CONTAINER_NAME=ghcr.io/nithinkatla00/rl-multileaderprotocolswitcher

.PHONY: proto
proto:
	@echo "Generating proto files..."
	@for target in $(PROTO_FILES) ; do \
		echo $$target ; \
		$(PROTOC) --plugin=protoc-gen-gogofaster=$(HOME)/go/bin/protoc-gen-gogofaster \
		--gogofaster_out=paths=source_relative,plugins=grpc:. -I. -I$(GOGO_PROTOBUF_PATH) $$target ; \
	done
	@echo "Done generating proto files."

TARGETS := peer client master erunner
TARGETDIR := bin

.PHONY: build
build:
	@echo "Go: Building..."
	@for target in $(TARGETS) ; do \
		$(GO) build -mod=vendor -o "$(TARGETDIR)/$$target" ./cmd/$$target || exit 1; \
	done
	@echo "Go: Build complete."

.PHONY: erunner
erunner: cmd/erunner/erunner.go
	$(GO) build -mod=vendor -o "$(TARGETDIR)/$@" $?

.PHONY: clean
clean:
	@$(RM) *.pb.*
	@$(RM) -r $(TARGETDIR)

container:
	@docker build -t $(REMOTE_CONTAINER_NAME) -f build/dockerfiles/Dockerfile .
	@docker tag $(REMOTE_CONTAINER_NAME) $(REMOTE_CONTAINER_NAME):$(DATE)
	@docker push $(REMOTE_CONTAINER_NAME):$(DATE)
	@docker push $(REMOTE_CONTAINER_NAME):latest