GO = go

FMTARGS = fmt
TESTARGS = test -cover -v
BUILDARGS = build -v

PROG = ./cgol_go

fmt:
	@$(GO) $(FMTARGS) ./...

test:
	@$(GO) $(TESTARGS) ./...

build:
	@$(GO) $(BUILDARGS)

run: build
	@$(PROG)

all:
	@clear
	@echo ""
	$(GO) $(FMTARGS) ./...
	@echo ""
	$(GO) $(TESTARGS) ./...
	@echo ""
	$(GO) $(BUILDARGS)
	@echo ""
	$(PROG)
	@echo ""

.PHONY: fmt test build run all
