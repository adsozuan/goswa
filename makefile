# HELP
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: run application in cmd/goswa
.PHONY: run
run:
	go run cmd/goswa/main.go 


## build: build application cmd/goswa into bin/
.PHONY: build
build:
	go build -ldflags='-s' -o=./bin/server ./cmd/goswa

## test: run tests
.PHONY: test
test:
	go test pkg/*