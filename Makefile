TAG := $(shell date +"%Y-%m-%d-%H-%M-%S")

build: lint
	go build -ldflags '-s -w' -o dist/brimstone main.go
dev:
	go run main.go
fmt:
	go fmt ./...
lint: fmt
	go vet ./...
start:
	./dist/brimstone
# Docker
build-docker:
	docker build . --tag impossible98/brimstone
	docker build . --tag impossible98/brimstone:$(TAG)	
push-docker: build-docker
	docker push impossible98/brimstone
	docker push impossible98/brimstone:$(TAG)
