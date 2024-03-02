SHELL := /bin/bash



.PHONY: build
build:
	go build .

.PHONY: run
run: 
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -t voterapi:latest -f Dockerfile .

.PHONY: docker-build-multi-builder
docker-build-multi-builder:
	docker buildx create --name=multi-builder --driver=docker-container
`
.PHONY: docker-build-multi-image
docker-build-multi-image:
	docker buildx build --platform linux/arm/v7,linux/arm64/8,linux/amd64 --builder multi-builder --push  -t maddchickenz/battaglialab_voterapi:latest .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./todo-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./todo-linux-arm64 .

