.PHONY: build run
.SILENT: build run

build:
	CGO_ENABLED=0 go build -o ./.bin/mockio ./cmd/main.go

run:
	./.bin/mockio -i ./examples/routes.json