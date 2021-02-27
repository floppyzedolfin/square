NAME=square
VERSION=0.1

build:
	CGO_ENABLED=0 go build -o build/square.out cmd/main.go
	docker build -t ${NAME}:{VERSION} build

run: stop
	docker run --rm -p 8530:3000 --name ${NAME} ${NAME}:${VERSION}

stagedbuild:
	docker build -t ${NAME}:${VERSION} .

stop:
	docker stop ${NAME} || true
