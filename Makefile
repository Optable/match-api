# Protobuf files management

PROTO_MATCH_FILES := $(wildcard match/v1/*.proto)
PROTO_FILES := $(PROTO_MATCH_FILES)

# Protobuf files and their generated equivalents
GO_MATCH_PB_FILES :=$(patsubst match/v1/%.proto,match/v1/%.pb.go,$(PROTO_MATCH_FILES))
GO_PB_FILES := $(GO_MATCH_PB_FILES)

.PHONY: build
build: clean
	protoc --proto_path . \
		--go_out=. --go_opt=paths=source_relative \
		$(PROTO_FILES)
	buf generate

.PHONY: clean
clean:
	rm -f $(GO_PB_FILES)

.PHONY: test
test:
	protoc --lint_out=sort_imports:. $(PROTO_FILES)

.PHONY: lint
lint:
	buf lint
