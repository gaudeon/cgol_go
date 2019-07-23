test:
	@go test -cover -v ./...

build:
	@go build

run: build
	@./cgol_go

all:
	@clear
	@echo ""
	go test -cover -v ./...
	@echo ""
	go build -v
	@echo ""
	./cgol_go
	@echo ""

.PHONY: test build run all
