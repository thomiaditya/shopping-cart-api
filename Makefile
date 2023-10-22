BINARY_NAME=main
GO_COMPILER=go
SOURCE_FILES=$(wildcard *.go)

# Set build target
build:
ifeq ($(shell uname),Darwin)
	@$(GO_COMPILER) build -o ./out/$(BINARY_NAME) ./cmd/server/$(SOURCE_FILES)
else ifeq ($(shell uname),Linux)
	@$(GO_COMPILER) build -o ./out/$(BINARY_NAME) ./cmd/server/$(SOURCE_FILES)
else
	@$(GO_COMPILER) build -o ./out/$(BINARY_NAME).exe ./cmd/server/$(SOURCE_FILES)
endif

# Run target
run: build
ifeq ($(shell uname),Windows_NT)
	./out/$(BINARY_NAME).exe
else
	./out/$(BINARY_NAME)
endif

# Set the clean target
clean:
ifeq ($(shell uname),Windows_NT)
	del /f ./out/$(BINARY_NAME).exe
else
	rm -f ./out/$(BINARY_NAME)
endif

# Set migration
migrate:
	@go run ./cmd/migrate/main.go