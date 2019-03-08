SRC_DIR := $(shell ls -d */|grep -vE 'vendor|dist|script|tmp')

.PHONY: all
all: help

## fmt: format/tidy go code
.PHONY: fmt
fmt:
	# gofmt code
	@gofmt -s -l -w $(SRC_DIR)

.PHONY: test
test:
	go test -v -coverprofile .cover.out ./...
	@go tool cover -func=.cover.out
	@go tool cover -html=.cover.out -o .cover.html

## test: test module example `make test/xxx`
.PHONY: test/%
test/%:
	go test -v -coverprofile ./$*/.cover.out ./$*
	go tool cover -func=./$*/.cover.out
	go tool cover -html=./$*/.cover.out -o ./$*/.cover.html

## clean: clean build
.PHONY: clean
clean:
	@find . -name '.cover.out' -type f | xargs rm -vf

.PHONY: help
help: Makefile
	@echo " Choose a command run in $(CMD):"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
