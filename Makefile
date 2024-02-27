.PHONY: help wire gen image run build test
.DEFAULT_GOAL := help

APP_NAME=blog-go
APP_PATH=.

# init
init:
	@go install github.com/google/wire/cmd/wire@latest

# wire init
wire:
	@cd cmd; wire

# gen, if wire_gen.go is not exist, run wire; if wire_gen.go is exist, run go generate
gen:
	@go generate ./...

# build image
image:
	@docker image build -t $(APP_NAME) .

# run
run:
	@go run $(APP_PATH)/main.go

# build
build:
	@mkdir -p bin && go mod tidy && go build -trimpath -ldflags="-w -s" -o bin/ $(APP_PATH)

# test
test:
	@go test -v ./...

# Show help
help:
	@echo ""
	@echo "Usage:"
	@echo "    make [target]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-_0-9]+:/ \
	{ \
		helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} { lastLine = $$0 }' $(MAKEFILE_LIST)
