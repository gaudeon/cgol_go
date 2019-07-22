test:
	@go test ./...

build:
	@go build

run: build
	@./cgol_go

all:
	@clear
	@echo ""
	go test ./...
	@echo ""
	go build -v
	@echo ""
	./cgol_go
	@echo ""

.PHONY: test build run all
