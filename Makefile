BINARY_NAME=main
GO_COMPILER=go
GO_FLAGS=-v
SOURCE_FILES=$(wildcard *.go)

# Set build target
build:
	@$(GO_COMPILER) build $(GO_FLAGS) -o ./out/$(BINARY_NAME) ./cmd/server/$(SOURCE_FILES)

# Run target
run: build
	@./out/$(BINARY_NAME)

# Set the clean target
clean:
	rm -f $(BINARY_NAME)
