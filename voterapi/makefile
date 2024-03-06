SHELL := /bin/bash

.PHONY: build
build:
	go build .

.PHONY: run
run: 
	go run main.go

.PHONY: run-tests
run-tests: 
	go clean -testcache
	go test -v ./...


.PHONY: get-voters
get-voters:
	curl -H "Content-Type: application/json" -X GET http://localhost:8080/voters/

.PHONY: populate
populate:
	curl -H "Content-Type: application/json" -X GET http://localhost:8080/voters/populate

