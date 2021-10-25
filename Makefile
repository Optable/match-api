# Protobuf files management

PROTO_EXTERNAL_FILES := $(wildcard external/*.proto)
PROTO_FILES := $(PROTO_EXTERNAL_FILES)

# Protobuf files and their generated equivalents
GO_EXTERNAL_PB_FILES :=$(patsubst external/%.proto,external/%.pb.go,$(PROTO_EXTERNAL_FILES))
GO_PB_FILES := $(GO_EXTERNAL_PB_FILES)

.PHONY: build
build: clean
	protoc --proto_path . \
		--go_out=. --go_opt=paths=source_relative \
		$(PROTO_FILES)

.PHONY: clean
clean:
	-rm -f $(GO_PB_FILES)

.PHONY: test
test:
	protoc --lint_out=sort_imports:. $(PROTO_FILES)