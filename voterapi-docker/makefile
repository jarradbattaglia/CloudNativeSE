SHELL := /bin/bash

.PHONY: build
build:
	go build .

.PHONY: run
run: 
	go run main.go

.PHONY: run-docker-compose
run-docker-compose: 
	docker-compose up

.PHONY: run-tests
run-tests: 
	go clean -testcache
	go test -v ./...

.PHONY: run-docker
run-docker: 
	docker run --rm -d --name cnse-redis -p 6379:6379 -p 8001:8001 redis/redis-stack:latest	
	docker run --rm -d --name jarrad-voterapi -p 8080:8080 -it voterapi:latest


.PHONY: cleanup-cache
cleanup-cache:
	go clean -testcache

.PHONY: cleanup
cleanup:
	docker kill cnse-redis 
	docker kill jarrad-voterapi


.PHONY: docker-build
docker-build:
	docker build -t voterapi:latest -f Dockerfile .

.PHONY: docker-build-multi-builder
docker-build-multi-builder:
	docker buildx create --name=multi-builder --driver=docker-container

.PHONY: docker-build-multi-image
docker-build-multi-image:
	docker buildx build --platform linux/arm/v7,linux/arm64/8,linux/amd64 --builder multi-builder --push  -t maddchickenz/battaglialab_voterapi:latest .

.PHONY: get-voters
get-voters:
	curl -H "Content-Type: application/json" -X GET http://localhost:8080/voters/

.PHONY: populate
populate:
	curl -H "Content-Type: application/json" -X GET http://localhost:8080/voters/populate