GO ?= go
PROTOC ?= protoc

GOGO_PROTOBUF_PATH := $(shell go list -e -f '{{.Dir}}' github.com/gogo/protobuf/gogoproto)/../
PROTO_FILES=$(shell find . -type f -not -path "./vendor/*" -name "*.proto")

.PHONY: proto
proto:
	@echo "Generating proto files..."
	@for target in $(PROTO_FILES) ; do \
		echo $$target ; \
		$(PROTOC) --gogofaster_out=paths=source_relative,plugins=grpc:. -I. -I$(GOGO_PROTOBUF_PATH) $$target ; \
	done
	@echo "Done generating proto files."

TARGETS := peer client master
TARGETDIR := bin

.PHONY: build
build:
	@echo "Go: Building..."
	@for target in $(TARGETS) ; do \
		$(GO) build -mod=vendor -o "$(TARGETDIR)/$$target" ./cmd/$$target || exit 1; \
	done
	@echo "Go: Build complete."

.PHONY: clean
clean:
	@$(RM) *.pb.*
	@$(RM) -r $(TARGETDIR)